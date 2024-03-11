package recordbestsingle

import (
	"encoding/json"
	"errors"
	"puzzle/app/middlewares/rabbitmq"
	"puzzle/app/models"
	commonService "puzzle/app/services"
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

func check(record *models.RecordBestSingle) error {
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
func Insert(record *models.RecordBestSingle) error {
	// 校验参数
	err := check(record)
	if err != nil {
		return err
	}

	record.RecordBreakCount = 1

	err = database.GetMySQL().Create(record).Error
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
func List(recordReq *models.RecordBestSingleReq) (models.RecordBestSingleListResp, error) {
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

	if recordReq.IdStr != "" {
		recordReq.Id, _ = strconv.ParseInt(recordReq.IdStr, 10, 64)
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
		db.Where("id = ?", recordReq.Id)
	}

	if recordReq.UserId != 0 {
		db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Dimension != 0 {
		db.Where("dimension = ?", recordReq.Dimension)
	}

	if recordReq.RecordId != 0 {
		db.Where("record_id = ?", recordReq.RecordId)
	}

	if len(recordReq.DurationRange) == 2 {
		if recordReq.DurationRange[0] != 0 {
			db.Where("record_duration >= ?", recordReq.DurationRange[0])
		}
		if recordReq.DurationRange[1] != 0 {
			db.Where("record_duration <= ?", recordReq.DurationRange[1])
		}
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
	err := db.Count(&recordListResp.Total).Error
	if err != nil {
		return recordListResp, errors.New("查询失败")
	}

	// 分页
	if recordReq.Pagination.Page > 0 && recordReq.Pagination.PageSize > 0 {
		db.Scopes(utils.Paginate(&recordReq.Pagination))
	}

	if recordReq.NeedUserInfo {
		db.Preload("UserInfo")
	}

	if recordReq.NeedRecordDetail {
		db.Preload("RecordDetail")
	}

	// 查询列表
	err = db.Find(&recordListResp.Records).Error
	if err != nil {
		return recordListResp, errors.New("查询失败")
	}

	return recordListResp, nil
}

// Update 更新记录
func Update(record *models.RecordBestSingle) error {
	db := database.GetMySQL().Table("record_best_single")

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
