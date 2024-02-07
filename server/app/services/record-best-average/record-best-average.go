package recordbestaverage

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
)

// check 校验参数
func check(record models.RecordBestAverage) error {
	if record.UserId == 0 {
		return errors.New("用户ID不能为空")
	}

	if record.Dimension == 0 {
		return errors.New("阶数不能为空")
	}

	if record.Type == 0 {
		return errors.New("类型不能为空")
	}

	if record.RecordIds == "" {
		return errors.New("记录ID不能为空")
	}

	if record.RecordAverageDuration == 0 {
		return errors.New("平均时长不能为空")
	}

	return nil
}

// Insert 插入一条记录
func Insert(record models.RecordBestAverage) error {
	// 校验参数
	err := check(record)
	if err != nil {
		return err
	}

	err = database.GetMySQL().Create(&record).Error
	if err != nil {
		return err
	}

	return nil
}

// List 记录列表
func List(recordReq models.RecordBestAverageReq) (models.RecordBestAverageListResp, error) {
	var recordListResp models.RecordBestAverageListResp
	db := database.GetMySQL().Table("record_best_average").Order("record_average_duration DESC")

	if recordReq.UserId != 0 {
		db = db.Where("user_id = ?", recordReq.UserId)
	}

	if recordReq.Dimension != 0 {
		db = db.Where("dimension = ?", recordReq.Dimension)
	}

	if recordReq.Type != 0 {
		db = db.Where("type = ?", recordReq.Type)
	}

	if len(recordReq.DurationRange) == 2 {
		db = db.Where("record_duration >= ? AND record_duration <= ?", recordReq.DurationRange[0], recordReq.DurationRange[1])
	}

	if len(recordReq.DateRange) == 2 {
		db = db.Where("created_at >= ? AND created_at <= ?", recordReq.DateRange[0], recordReq.DateRange[1])
	}

	// 查询总数
	err := db.Count(&recordListResp.Total).Error
	if err != nil {
		return recordListResp, errors.New("记录总数查询失败")
	}

	// 分页
	if recordReq.Pagination.Page > 0 && recordReq.Pagination.PageSize > 0 {
		db = db.Scopes(utils.Paginate(&recordReq.Pagination))
	}

	err = db.Find(&recordListResp.Records).Error
	if err != nil {
		return recordListResp, errors.New("记录查询失败")
	}

	return recordListResp, nil
}

// Update 更新记录
func Update(record models.RecordBestAverage) error {
	db := database.GetMySQL().Table("record_best_average").Where("user_id = ? AND dimension = ? AND type = ?", record.UserId, record.Dimension, record.Type)

	if record.RecordIds != "" {
		db = db.Where("record_ids = ?", record.RecordIds)
	}

	if record.RecordAverageDuration != 0 {
		db = db.Where("record_average_duration = ?", record.RecordAverageDuration)
	}

	err := db.Updates(&record).Error
	if err != nil {
		return err
	}

	return nil
}
