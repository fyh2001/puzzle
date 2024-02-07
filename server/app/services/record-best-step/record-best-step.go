package recordbeststep

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
)

// check 检查参数
func check(record models.RecordBestStep) error {
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
func Insert(record models.RecordBestStep) error {
	err := check(record)
	if err != nil {
		return err
	}

	err = database.GetMySQL().Create(&record).Error
	if err != nil {
		return errors.New("添加失败")
	}

	return nil
}

// List 查询记录
func List(recordReq models.RecordBestStepReq) (models.RecordBestStepListResp, error) {
	var recordBestStepListResp models.RecordBestStepListResp
	db := database.GetMySQL().Table("record_best_step").Order("record_step DESC")

	if recordReq.UserId != 0 {
		db = db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Dimension != 0 {
		db = db.Where("dimension = ?", recordReq.Dimension)
	}

	if recordReq.RecordId != 0 {
		db = db.Where("record_id = ?", recordReq.RecordId)
	}

	if len(recordReq.StepRange) == 2 {
		db = db.Where("record_step >= ? AND record_step <= ?", recordReq.StepRange[0], recordReq.StepRange[1])
	}

	if len(recordReq.DateRange) == 2 {
		db = db.Where("created_at >= ? AND created_at <= ?", recordReq.DateRange[0], recordReq.DateRange[1])
	}

	// 查询总数
	err := db.Count(&recordBestStepListResp.Total).Error
	if err != nil {
		return recordBestStepListResp, errors.New("总数查询失败")
	}

	// 分页
	if (recordReq.Pagination.Page > 0) && (recordReq.Pagination.PageSize > 0) {
		db = db.Scopes(utils.Paginate(&recordReq.Pagination))
	}

	// 查询记录
	err = db.Find(&recordBestStepListResp.Records).Error
	if err != nil {
		return recordBestStepListResp, errors.New("查询失败")
	}

	return recordBestStepListResp, nil
}

// Update 更新记录
func Update(record models.RecordBestStep) error {
	db := database.GetMySQL().Table("record_best_step").Where("user_id = ? AND dimension = ?", record.UserId, record.Dimension)

	if record.RecordId != 0 {
		db = db.Where("record_id = ?", record.RecordId)
	}

	if record.RecordStep != 0 {
		db = db.Where("record_step = ?", record.RecordStep)
	}

	err := db.Updates(&record).Error

	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}
