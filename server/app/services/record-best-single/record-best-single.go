package recordbestsingle

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
	mq := rabbitmq.NewRabbitMQ("best_single_rank_update_queue", "", "")
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

func check(record models.RecordBestSingle) error {
	if record.UserId == 0 {
		return errors.New("用户ID不能为空")
	}

	if record.Dimension == 0 {
		return errors.New("阶数不能为空")
	}

	if record.RecordDuration == 0 {
		return errors.New("耗时不能为空")
	}

	if record.RecordId == 0 {
		return errors.New("记录ID不能为空")
	}

	if record.RecordStep == 0 {
		return errors.New("步数不能为空")
	}

	return nil
}

// Insert 插入一条记录
func Insert(record models.RecordBestSingle) error {
	// 校验参数
	err := check(record)
	if err != nil {
		return err
	}

	record.RecordBreakCount = 1

	err = database.GetMySQL().Create(&record).Error
	if err != nil {
		return err
	}

	// 发送消息至消息队列
	publishMessage(commonService.RankUpdate{
		Dimension: record.Dimension,
	})

	return nil
}

// List 查询记录列表
func List(recordReq models.RecordBestSingleReq) (models.RecordBestSingleListResp, error) {
	var recordListResp models.RecordBestSingleListResp

	if recordReq.Username != "" || recordReq.Nickname != "" {
		userInfo, err := commonService.GetUserInfoByUsernameOrNickname(recordReq.Username, recordReq.Nickname)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return recordListResp, errors.New("查询用户信息失败")
		}

		if userInfo.Id == "" {
			recordReq.UserIdStr = "-1"
		} else {
			recordReq.UserIdStr = userInfo.Id
		}
	}

	if recordReq.RecordIdStr != "" {
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

	db := database.GetMySQL().Table("record_best_single").Order(recordReq.OrderBy + " " + recordReq.Sorted)

	if recordReq.Id != 0 {
		db = db.Where("id = ?", recordReq.Id)
	}

	if recordReq.UserId != 0 {
		db = db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Dimension != 0 {
		db = db.Where("dimension = ?", recordReq.Dimension)
	}

	if recordReq.RecordId != 0 {
		db = db.Where("record_id = ?", recordReq.RecordId)
	}

	if len(recordReq.DurationRange) == 2 {
		if recordReq.DurationRange[0] != 0 {
			db = db.Where("record_duration >= ?", recordReq.DurationRange[0])
		}
		if recordReq.DurationRange[1] != 0 {
			db = db.Where("record_duration <= ?", recordReq.DurationRange[1])
		}
	}

	if len(recordReq.StepRange) == 2 {
		if recordReq.StepRange[0] != 0 {
			db = db.Where("record_step >= ?", recordReq.StepRange[0])
		}
		if recordReq.StepRange[1] != 0 {
			db = db.Where("record_step <= ?", recordReq.StepRange[1])
		}
	}

	if len(recordReq.RankRange) == 2 {
		if recordReq.RankRange[0] != 0 {
			db = db.Where("ranked >= ?", recordReq.RankRange[0])
		}
		if recordReq.RankRange[01] != 0 {
			db = db.Where("ranked <= ?", recordReq.RankRange[1])
		}
	}

	if len(recordReq.BreakCountRange) == 2 {
		if recordReq.BreakCountRange[0] != 0 {
			db = db.Where("record_break_count >= ?", recordReq.BreakCountRange[0])
		}
		if recordReq.BreakCountRange[1] != 0 {
			db = db.Where("record_break_count <= ?", recordReq.BreakCountRange[1])
		}
	}

	if len(recordReq.DateRange) == 2 && !recordReq.DateRange[0].IsZero() && !recordReq.DateRange[1].IsZero() {
		db = db.Where("created_at >= ? AND created_at <= ?", recordReq.DateRange[0], recordReq.DateRange[1])
	}

	// 查询总数
	err := db.Count(&recordListResp.Total).Error
	if err != nil {
		return recordListResp, errors.New("查询失败")
	}

	// 分页
	if recordReq.Pagination.Page > 0 && recordReq.Pagination.PageSize > 0 {
		db = db.Scopes(utils.Paginate(&recordReq.Pagination))
	}

	// 查询列表
	err = db.Find(&recordListResp.Records).Error
	if err != nil {
		return recordListResp, errors.New("查询失败")
	}

	if recordReq.NeedUserInfo {
		// 查询用户信息
		userIds := make([]int64, 0)
		for _, record := range recordListResp.Records {
			userId, _ := strconv.ParseInt(record.UserId, 10, 64)
			userIds = append(userIds, userId)
		}

		userList, err := commonService.GetUserInfo(userIds)
		if err != nil {
			return recordListResp, errors.New("查询用户信息失败")
		}

		userMap := make(map[string]models.UserResp)
		for _, user := range userList.Records {
			userMap[user.Id] = user
		}

		for i, record := range recordListResp.Records {
			recordListResp.Records[i].UserInfo = userMap[record.UserId]
		}
	}

	if recordReq.NeedRecordDetail {
		// 查询记录详情
		recordIds := make([]int64, 0)
		for _, record := range recordListResp.Records {
			recordId, _ := strconv.ParseInt(record.RecordId, 10, 64)
			recordIds = append(recordIds, recordId)
		}

		recordList, err := commonService.GetRecordDetail(recordIds)
		if err != nil {
			return recordListResp, errors.New("查询记录详情失败")
		}

		recordMap := make(map[string][]models.RecordResp)

		for _, record := range recordList.Records {
			key := record.UserId + "-" + strconv.Itoa(record.Dimension)

			recordMap[key] = append(recordMap[key], record)
		}

		for i, record := range recordListResp.Records {

			key := record.UserId + "-" + strconv.Itoa(record.Dimension)

			recordListResp.Records[i].RecordDetail.Records = recordMap[key]
			recordListResp.Records[i].RecordDetail.Total = int64(len(recordMap[key]))
		}
	}

	return recordListResp, nil
}

// ListWithUserInfo 查询记录列表并携带用户信息
func ListWithUserInfo(recordReq models.RecordBestSingleReq) (models.RecordBestSingleListResp, error) {
	var recordListResp models.RecordBestSingleListResp
	db := database.GetMySQL().Table("record_best_single").Order("record_duration " + recordReq.Sorted)

	if recordReq.UserId != 0 {
		db = db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Dimension != 0 {
		db = db.Where("dimension = ?", recordReq.Dimension)
	}

	if recordReq.RecordId != 0 {
		db = db.Where("record_id = ?", recordReq.RecordId)
	}

	if len(recordReq.DurationRange) == 2 {
		db = db.Where("record_duration >= ? AND record_duration <= ?", recordReq.DurationRange[0], recordReq.DurationRange[1])
	}

	if len(recordReq.StepRange) == 2 {
		db = db.Where("record_step >= ? AND record_step <= ?", recordReq.StepRange[0], recordReq.StepRange[1])
	}

	if len(recordReq.DateRange) == 2 {
		db = db.Where("created_at >= ? AND created_at <= ?", recordReq.DateRange[0], recordReq.DateRange[1])
	}

	// 查询总数
	err := db.Count(&recordListResp.Total).Error
	if err != nil {
		return recordListResp, errors.New("查询失败")
	}

	// 分页
	if recordReq.Pagination.Page > 0 && recordReq.Pagination.PageSize > 0 {
		db = db.Scopes(utils.Paginate(&recordReq.Pagination))
	}

	// 查询列表
	err = db.Find(&recordListResp.Records).Error
	if err != nil {
		return recordListResp, errors.New("查询失败")
	}

	// 查询用户信息
	userIds := make([]int64, 0)
	for _, record := range recordListResp.Records {
		userId, _ := strconv.ParseInt(record.UserId, 10, 64)
		userIds = append(userIds, userId)
	}

	userReq := models.UserReq{
		Ids: userIds,
	}

	userList, err := userService.List(userReq)
	if err != nil {
		return recordListResp, errors.New("查询用户信息失败")
	}

	userMap := make(map[string]models.UserResp)
	for _, user := range userList.Records {
		userMap[user.Id] = user
	}

	for i, record := range recordListResp.Records {
		recordListResp.Records[i].UserInfo = userMap[record.UserId]
	}

	return recordListResp, nil
}

// Update 更新记录
func Update(record models.RecordBestSingle) error {
	db := database.GetMySQL().Table("record_best_single")

	err := db.Updates(&record).Error

	if err != nil {
		return errors.New("更新失败")
	}

	// 发送消息至消息队列
	publishMessage(commonService.RankUpdate{
		Dimension: record.Dimension,
	})

	return nil
}
