package record

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
)

// Insert 新增记录
func Insert(record models.Record) error {
	err := database.GetMySQL().Create(&record).Error
	if err != nil {
		return errors.New("新增失败")
	}

	return nil
}

// List 记录列表
func List(recordReq models.RecordReq) (models.RecordListResp, error) {
	var recordListResp models.RecordListResp
	db := database.GetMySQL().Table("record")

	if recordReq.Id != 0 {
		db = db.Where("id = ?", recordReq.Id)
	}

	if recordReq.UserId != 0 {
		db = db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Event != "" {
		db = db.Where("event = ?", recordReq.Event)
	}

	if recordReq.Type != 0 {
		db = db.Where("type = ?", recordReq.Type)
	}

	if len(recordReq.DurationRange) == 2 {
		db = db.Where("duration >= ? AND duration <= ?", recordReq.DurationRange[0], recordReq.DurationRange[1])
	}

	if len(recordReq.StepRange) == 2 {
		db = db.Where("step >= ? AND step <= ?", recordReq.StepRange[0], recordReq.StepRange[1])
	}

	if recordReq.Status != 0 {
		db = db.Where("status = ?", recordReq.Status)
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

	// 查询记录
	err = db.Find(&recordListResp.Records).Error
	if err != nil {
		return recordListResp, errors.New("查询失败")
	}

	return recordListResp, nil
}

// Update 更新记录
func Update(record models.Record) error {
	err := database.GetMySQL().Model(&record).Updates(&record).Error
	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}
