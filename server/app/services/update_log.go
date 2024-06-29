package services

import (
	"encoding/json"
	"errors"
	"puzzle/app/middlewares/rabbitmq"
	"puzzle/app/middlewares/rabbitmq/handlers"
	"puzzle/app/middlewares/websocket"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
	"strconv"
)

type UpdateLogService interface {
	check(updateLog *models.UpdateLog) error
	Insert(updateLog *models.UpdateLog) error
	List(updateLogReq *models.UpdateLogReq) (models.UpdateLogListResp, error)
}

type UpdateLogImpl struct {
}

// publishMessage 发送消息至消息队列
func (UpdateLogImpl) publishMessage(notificationMsg handlers.NotificationMsg) {
	mq := rabbitmq.NewRabbitMQ("notification_queue", "", "")
	defer mq.Destory()

	message := rabbitmq.RabbitMQMessage{
		NotificationMsg: notificationMsg,
		Message:         "notification-all",
	}

	messageByte, err := json.Marshal(message)
	if err != nil {
		return
	}

	mq.Publish(messageByte)
}

// sendWebsocketMessage 发送消息至websocket
func (UpdateLogImpl) sendWebsocketMessage(userId int64, content string) {

	// 编辑消息
	message, _ := json.Marshal(websocket.Message{
		Type:    "update",
		Content: content,
	})

	// 全体消息
	if userId == 0 {
		websocket.ClientManagerInstance.Broadcast <- message
	} else { // 单体消息
		websocketUser, ok := websocket.GatewayUser.Load(strconv.FormatInt(userId, 10))
		if ok {
			userBase := websocketUser.(*websocket.WebSocketUser)
			for _, v := range userBase.ClientID {
				client, ok := websocket.ClientManagerInstance.Clients.Load(v)
				if ok {
					clientInfo := client.(*websocket.Client)
					clientInfo.Send <- message
				}
			}
		}
	}
}

func (UpdateLogImpl) check(updateLog *models.UpdateLog) error {

	if updateLog.Version == "" {
		return errors.New("[update_log]版本号不能为空")
	}

	if updateLog.Type == 0 {
		return errors.New("[update_log]更新类型不能为空")
	}

	if updateLog.Content == "" {
		return errors.New("[update_log]更新内容不能为空")
	}

	return nil
}

func (UpdateLogImpl) Insert(updateLog *models.UpdateLog) error {
	err := UpdateLog.check(updateLog)
	if err != nil {
		return err
	}
	snowflake := utils.Snowflake{}

	updateLog.Id = snowflake.NextVal() // 生成ID
	updateLog.Status = 1

	err = database.GetMySQL().Create(updateLog).Error
	if err != nil {
		return errors.New("[update_log]新增更新日志失败")
	}

	userIds, err := User.GetAllUserId()
	if err != nil {
		return errors.New("[update_log]获取用户ID失败")
	}

	// 发送消息
	notificationMsg := handlers.NotificationMsg{
		TypeId:  1,
		UserIds: userIds,
		Content: "当前版本已更新为" + updateLog.Version + ", 可前往<更新日志>查看更新详情．",
	}
	UpdateLog.publishMessage(notificationMsg)

	// 发送消息
	UpdateLog.sendWebsocketMessage(0, "即将发布新版本, 为防止数据丢失, 请在30秒内下线")
	return nil
}

func (UpdateLogImpl) List(updateLogReq *models.UpdateLogReq) (models.UpdateLogListResp, error) {
	var updateLogListResp models.UpdateLogListResp

	if updateLogReq.IdStr != "" {
		updateLogReq.Id, _ = strconv.ParseInt(updateLogReq.IdStr, 10, 64)
	}

	if updateLogReq.OrderBy == "" {
		updateLogReq.OrderBy = "id"
	}

	db := database.GetMySQL().Model(&models.UpdateLog{}).Order(updateLogReq.OrderBy + " " + updateLogReq.Sorted)

	if updateLogReq.Id != 0 {
		db = db.Where("id = ?", updateLogReq.Id)
	}

	if updateLogReq.Version != "" {
		db = db.Where("verson = ?", updateLogReq.Version)
	}

	if updateLogReq.Status != 0 {
		db = db.Where("status = ?", updateLogReq.Status)
	}

	err := db.Count(&updateLogListResp.Total).Error
	if err != nil {
		return updateLogListResp, errors.New("[update_log]查询更新日志总数失败")
	}

	if len(updateLogReq.DateRange) == 2 && !updateLogReq.DateRange[0].IsZero() && !updateLogReq.DateRange[1].IsZero() {
		db.Where("created_at >= ? AND created_at <= ?", updateLogReq.DateRange[0], updateLogReq.DateRange[1])
	}

	err = db.Find(&updateLogListResp.Records).Error
	if err != nil {
		return updateLogListResp, errors.New("[update_log]查询更新日志失败")
	}

	return updateLogListResp, nil
}
