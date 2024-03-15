package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"puzzle/app/middlewares/rabbitmq"
	"puzzle/app/middlewares/rabbitmq/handlers"
	"puzzle/app/middlewares/websocket"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

var currentTitle = "notification"

type NotificationService interface {
	check(notification *models.Notification) error
	Insert(notificationReq *models.NotificationReq) error
	List(notificationReq *models.NotificationReq) (models.NotificationListResp, error)
	Update(notificationReq *models.NotificationReq) error
	ReadAll(notificationReq *models.NotificationReq) error
}

type NotificationImpl struct{}

// publishMessage 发送消息至消息队列
func (NotificationImpl) publishMessage(notificationMsg handlers.NotificationMsg) {
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
func (NotificationImpl) sendWebsocketMessage(userId int64, content string) {

	// 编辑消息
	message, _ := json.Marshal(websocket.Message{
		Type:    "notification",
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

func (NotificationImpl) check(notification *models.Notification) error {
	return nil
}

func (NotificationImpl) Insert(notificationReq *models.NotificationReq) error {

	if notificationReq.UserIdStr != "" {
		notificationReq.UserId, _ = strconv.ParseInt(notificationReq.UserIdStr, 10, 64)
	}

	snowflake := utils.Snowflake{}

	// 发送消息
	Notification.sendWebsocketMessage(notificationReq.UserId, notificationReq.Content)

	// 如果是全体通知
	if notificationReq.UserId == 0 {
		userIds, err := User.GetAllUserId()
		if err != nil {
			return fmt.Errorf("[%s]获取用户id失败", currentTitle)
		}

		notificationMsg := handlers.NotificationMsg{
			TypeId:  notificationReq.TypeId,
			Content: notificationReq.Content,
			UserIds: userIds,
		}

		Notification.publishMessage(notificationMsg)

		return nil
	}

	notification := &models.Notification{
		Id:         snowflake.NextVal(),
		UserId:     notificationReq.UserId,
		TypeId:     notificationReq.TypeId,
		Content:    notificationReq.Content,
		ReadStatus: 1,
		Status:     1,
	}

	err := Notification.check(notification)
	if err != nil {
		return err
	}

	err = database.GetMySQL().Create(notification).Error
	if err != nil {
		return err
	}

	return nil
}

func (NotificationImpl) List(notificationReq *models.NotificationReq) (models.NotificationListResp, error) {
	var notificationListResp models.NotificationListResp

	if notificationReq.Username != "" || notificationReq.Nickname != "" {
		userInfo, err := User.GetUserByUsernameOrNickname(notificationReq.Username, notificationReq.Nickname)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return notificationListResp, errors.New("查询用户信息失败")
		}

		if userInfo.Id == "" {
			notificationReq.UserIdStr = "-1"
		} else {
			notificationReq.UserIdStr = userInfo.Id
		}
	}

	if notificationReq.IdStr != "" {
		notificationReq.Id, _ = strconv.ParseInt(notificationReq.IdStr, 10, 64)
	}

	if notificationReq.UserIdStr != "" {
		notificationReq.UserId, _ = strconv.ParseInt(notificationReq.UserIdStr, 10, 64)
	}

	if len(notificationReq.IdsStr) > 0 {
		notificationReq.Ids = make([]int64, len(notificationReq.IdsStr))
		for i, v := range notificationReq.IdsStr {
			notificationReq.Ids[i], _ = strconv.ParseInt(v, 10, 64)
		}
	}

	if notificationReq.OrderBy == "" {
		notificationReq.OrderBy = "updated_at"
	}

	db := database.GetMySQL().Table("notification").Order("read_status asc, created_at desc")

	if notificationReq.Sorted != "" {
		db.Order(notificationReq.OrderBy + " " + notificationReq.Sorted)
	}

	if notificationReq.Id != 0 {
		db.Where("id = ?", notificationReq.Id)
	}

	if len(notificationReq.Ids) > 0 {
		db.Where("id in ?", notificationReq.Ids)
	}

	if notificationReq.UserId != 0 {
		db.Where("user_id = ?", notificationReq.UserId)
	}

	if notificationReq.TypeId != 0 {
		db.Where("type_id = ?", notificationReq.TypeId)
	}

	if notificationReq.Content != "" {
		db.Where("content Like ?", "%"+notificationReq.Content+"%")
	}

	if notificationReq.ReadStatus != 0 {
		db.Where("read_status = ?", notificationReq.ReadStatus)
	}

	if notificationReq.Status != 0 {
		db.Where("status = ?", notificationReq.Status)
	}

	if len(notificationReq.DateRange) == 2 {
		if notificationReq.DateRange[0] != (time.Time{}) {
			db.Where("created_at >= ?", notificationReq.DateRange[0])
		}
		if notificationReq.DateRange[1] != (time.Time{}) {
			db.Where("created_at <= ?", notificationReq.DateRange[1])
		}
	}

	// 查询总数
	err := db.Count(&notificationListResp.Total).Error
	if err != nil {
		return notificationListResp, errors.New("查询通知总数失败")
	}

	if notificationReq.OnlyTotal {
		return notificationListResp, nil
	}

	db.Preload("NotificationTypeInfo")

	// 分页
	if notificationReq.Pagination.Page > 0 && notificationReq.Pagination.PageSize > 0 {
		db.Scopes(utils.Paginate(&notificationReq.Pagination))
	}

	// 查询列表
	err = db.Find(&notificationListResp.Records).Error
	if err != nil {
		return notificationListResp, errors.New(("查询通知列表失败"))
	}

	return notificationListResp, nil
}

func (NotificationImpl) Update(notificationReq *models.NotificationReq) error {

	if notificationReq.IdStr != "" {
		notificationReq.Id, _ = strconv.ParseInt(notificationReq.IdStr, 10, 64)
	}

	if notificationReq.UserIdStr != "" {
		notificationReq.UserId, _ = strconv.ParseInt(notificationReq.UserIdStr, 10, 64)
	}

	if len(notificationReq.IdsStr) > 0 {
		notificationReq.Ids = make([]int64, len(notificationReq.IdsStr))
		for i, v := range notificationReq.IdsStr {
			notificationReq.Ids[i], _ = strconv.ParseInt(v, 10, 64)
		}
	}

	notification := &models.Notification{
		Id:         notificationReq.Id,
		UserId:     notificationReq.UserId,
		TypeId:     notificationReq.TypeId,
		Content:    notificationReq.Content,
		ReadStatus: notificationReq.ReadStatus,
		Status:     notificationReq.Status,
	}

	err := Notification.check(notification)
	if err != nil {
		return err
	}

	err = database.GetMySQL().Model(&models.Notification{}).Where("id = ?", notificationReq.Id).Updates(notification).Error
	if err != nil {
		return err
	}

	return nil
}

// 一键已读
func (NotificationImpl) ReadAll(notificationReq *models.NotificationReq) error {
	if notificationReq.UserIdStr != "" {
		notificationReq.UserId, _ = strconv.ParseInt(notificationReq.UserIdStr, 10, 64)
	}

	notification := &models.Notification{
		UserId:     notificationReq.UserId,
		ReadStatus: 2,
	}

	err := database.GetMySQL().Model(&models.Notification{}).Where("user_id = ?", notificationReq.UserId).Updates(notification).Error
	if err != nil {
		return fmt.Errorf("[%s]一键已读失败", currentTitle)
	}

	return nil
}
