package recordbeststep

import (
	"encoding/json"
	"errors"
	"puzzle/app/middlewares/rabbitmq"
	"puzzle/app/models"
	commonService "puzzle/app/services"
	userService "puzzle/app/services/user"
	"puzzle/database"
	"puzzle/utils"
	"strconv"

	"gorm.io/gorm"
)

// publishMessage 发送消息至消息队列
func publishMessage(rankUpdate commonService.RankUpdate) {
	mq := rabbitmq.NewRabbitMQ("best_step_rank_update_queue", "", "")
	defer mq.Destory()

	message := rabbitmq.RabbitMQMessage{
		RankUpdate: rankUpdate,
		Message:    "rank update",
	}

	messageByte, err := json.Marshal(message)
	if err != nil {
		return
	}

	mq.Publish(messageByte)
}

// check 检查参数
func check(record *models.RecordBestStep) error {
	if record.UserId == 0 {
		return errors.New("用户ID不能为空")
	}

	if record.Dimension == 0 {
		return errors.New("阶数不能为空")
	}

	if record.RecordId == 0 {
		return errors.New("记录ID不能为空")
	}

	if record.RecordStep == 0 {
		return errors.New("步数不能为空")
	}

	return nil
}

// Insert 添加记录
func Insert(record *models.RecordBestStep) error {
	err := check(record)
	if err != nil {
		return err
	}

	err = database.GetMySQL().Create(record).Error
	if err != nil {
		return errors.New("添加失败")
	}

	// 发送消息至消息队列
	publishMessage(commonService.RankUpdate{
		Dimension: record.Dimension,
	})

	return nil
}

// List 查询记录
func List(recordReq *models.RecordBestStepReq) (models.RecordBestStepListResp, error) {
	var recordBestStepListResp models.RecordBestStepListResp

	if recordReq.Username != "" || recordReq.Nickname != "" {
		userInfo, err := commonService.GetUserInfoByUsernameOrNickname(recordReq.Username, recordReq.Nickname)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return recordBestStepListResp, errors.New("查询用户信息失败")
		}

		if userInfo.Id == "" {
			recordReq.UserIdStr = "-1"
		} else {
			recordReq.UserIdStr = userInfo.Id
		}
	}

	if recordReq.IdStr != "" {
		recordReq.Id, _ = strconv.ParseInt(recordReq.IdStr, 10, 64)
	}
	if recordReq.UserIdStr != "" {
		recordReq.UserId, _ = strconv.ParseInt(recordReq.UserIdStr, 10, 64)
	}

	if recordReq.RecordIdStr != "" {
		recordReq.RecordId, _ = strconv.ParseInt(recordReq.RecordIdStr, 10, 64)
	}

	if recordReq.OrderBy == "" {
		recordReq.OrderBy = "id"
	}

	db := database.GetMySQL().Table("record_best_step").Order(recordReq.OrderBy + " " + recordReq.Sorted)

	if recordReq.UserId != 0 {
		db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Dimension != 0 {
		db.Where("dimension = ?", recordReq.Dimension)
	}

	if recordReq.RecordId != 0 {
		db.Where("record_id = ?", recordReq.RecordId)
	}

	if len(recordReq.StepRange) == 2 {
		if recordReq.StepRange[0] != 0 {
			db.Where("record_step >= ?", recordReq.StepRange[0])
		}
		if recordReq.StepRange[1] != 0 {
			db.Where("record_step <= ?", recordReq.StepRange[1])
		}
	}

	if len(recordReq.RankRange) == 2 {
		if recordReq.RankRange[0] != 0 {
			db.Where("ranked >= ?", recordReq.RankRange[0])
		}
		if recordReq.RankRange[01] != 0 {
			db.Where("ranked <= ?", recordReq.RankRange[1])
		}
	}

	if len(recordReq.BreakCountRange) == 2 {
		if recordReq.BreakCountRange[0] != 0 {
			db.Where("record_break_count >= ?", recordReq.BreakCountRange[0])
		}
		if recordReq.BreakCountRange[1] != 0 {
			db.Where("record_break_count <= ?", recordReq.BreakCountRange[1])
		}
	}

	if len(recordReq.DateRange) == 2 && !recordReq.DateRange[0].IsZero() && !recordReq.DateRange[1].IsZero() {
		db.Where("created_at >= ? AND created_at <= ?", recordReq.DateRange[0], recordReq.DateRange[1])
	}

	// 查询总数
	err := db.Count(&recordBestStepListResp.Total).Error
	if err != nil {
		return recordBestStepListResp, errors.New("总数查询失败")
	}

	// 分页
	if (recordReq.Pagination.Page > 0) && (recordReq.Pagination.PageSize > 0) {
		db.Scopes(utils.Paginate(&recordReq.Pagination))
	}

	if recordReq.NeedUserInfo {
		db.Preload("UserInfo")
	}

	// 查询记录
	err = db.Find(&recordBestStepListResp.Records).Error
	if err != nil {
		return recordBestStepListResp, errors.New("查询失败")
	}

	if recordReq.NeedRecordDetail {
		// 查询记录详情
		recordIds := make([]int64, 0)
		for _, record := range recordBestStepListResp.Records {
			recordId, _ := strconv.ParseInt(record.RecordId, 10, 64)
			recordIds = append(recordIds, recordId)
		}

		recordList, err := commonService.GetRecordDetail(recordIds)
		if err != nil {
			return recordBestStepListResp, errors.New("查询记录详情失败")
		}

		recordMap := make(map[string][]models.RecordResp)

		for _, record := range recordList.Records {
			recordMap[record.UserId] = append(recordMap[record.UserId], record)
		}

		for i, record := range recordBestStepListResp.Records {
			recordBestStepListResp.Records[i].RecordDetail.Records = recordMap[record.UserId]
			recordBestStepListResp.Records[i].RecordDetail.Total = int64(len(recordMap[record.UserId]))
		}
	}

	return recordBestStepListResp, nil
}

// ListWithUserInfo 查询记录并携带用户信息
func ListWithUserInfo(recordReq *models.RecordBestStepReq) (models.RecordBestStepListResp, error) {
	var recordBestStepListResp models.RecordBestStepListResp
	db := database.GetMySQL().Table("record_best_step").Order("record_step " + recordReq.Sorted)

	if recordReq.UserId != 0 {
		db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Dimension != 0 {
		db.Where("dimension = ?", recordReq.Dimension)
	}

	if recordReq.RecordId != 0 {
		db.Where("record_id = ?", recordReq.RecordId)
	}

	if len(recordReq.StepRange) == 2 {
		db.Where("record_step >= ? AND record_step <= ?", recordReq.StepRange[0], recordReq.StepRange[1])
	}

	if len(recordReq.DateRange) == 2 {
		db.Where("created_at >= ? AND created_at <= ?", recordReq.DateRange[0], recordReq.DateRange[1])
	}

	// 查询总数
	err := db.Count(&recordBestStepListResp.Total).Error
	if err != nil {
		return recordBestStepListResp, errors.New("总数查询失败")
	}

	// 分页
	if (recordReq.Pagination.Page > 0) && (recordReq.Pagination.PageSize > 0) {
		db.Scopes(utils.Paginate(&recordReq.Pagination))
	}

	// 查询记录
	err = db.Find(&recordBestStepListResp.Records).Error
	if err != nil {
		return recordBestStepListResp, errors.New("查询失败")
	}

	// 查询用户信息
	userIds := make([]int64, 0)
	for _, record := range recordBestStepListResp.Records {
		userId, _ := strconv.ParseInt(record.UserId, 10, 64)
		userIds = append(userIds, userId)
	}

	userReq := models.UserReq{
		Ids: userIds,
	}

	userList, err := userService.List(&userReq)
	if err != nil {
		return recordBestStepListResp, errors.New("查询用户信息失败")
	}

	userMap := make(map[string]models.UserResp)
	for _, user := range userList.Records {
		userMap[user.Id] = user
	}

	for i, record := range recordBestStepListResp.Records {
		recordBestStepListResp.Records[i].UserInfo = userMap[record.UserId]
	}

	return recordBestStepListResp, nil
}

// Update 更新记录
func Update(record *models.RecordBestStep) error {
	db := database.GetMySQL().Table("record_best_step").Where("user_id = ? AND dimension = ?", record.UserId, record.Dimension)

	err := db.Updates(record).Error

	if err != nil {
		return errors.New("更新失败")
	}

	// 发送消息至消息队列
	publishMessage(commonService.RankUpdate{
		Dimension: record.Dimension,
	})

	return nil
}
