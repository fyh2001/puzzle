package recordbestsingle

import (
	"errors"
	"puzzle/app/models"
	userService "puzzle/app/services/user"
	"puzzle/database"
	"puzzle/utils"
	"strconv"
)

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

	err = database.GetMySQL().Create(&record).Error
	if err != nil {
		return err
	}

	return nil
}

// List 查询记录列表
func List(recordReq models.RecordBestSingleReq) (models.RecordBestSingleListResp, error) {
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
	db := database.GetMySQL().Table("record_best_single").Where("user_id = ? AND dimension = ?", record.UserId, record.Dimension)

	if record.RecordId != 0 {
		db = db.Where("record_id = ?", record.RecordId)
	}

	if record.RecordDuration != 0 {
		db = db.Where("record_duration = ?", record.RecordDuration)
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
