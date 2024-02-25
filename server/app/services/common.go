package services

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
)

func GetUserInfoByUsernameOrNickname(username string, nickname string) (models.UserResp, error) {
	var user models.UserResp

	db := database.GetMySQL().Table("user")

	if username != "" {
		db = db.Where("username Like ?", "%"+username+"%")
	}

	if nickname != "" {
		db = db.Where("nickname Like ?", "%"+nickname+"%")
	}

	err := db.First(&user).Error
	return user, err
}

// GetUserInfo 获取用户信息
func GetUserInfo(userIds []int64) (models.UserListResp, error) {
	var userInfo models.UserListResp

	err := database.GetMySQL().Table("user").Where("id in ?", userIds).Find(&userInfo.Records).Error
	if err != nil {
		return userInfo, err
	}

	return userInfo, nil
}

// GetRecordDetail 获取记录详情
func GetRecordDetail(recordIds []int64) (models.RecordListResp, error) {
	var recordDetail models.RecordListResp

	err := database.GetMySQL().Table("record").Where("id in ?", recordIds).Find(&recordDetail.Records).Error
	if err != nil {
		return recordDetail, err
	}

	return recordDetail, nil
}

// UpdateRecordBestSingleRank 更新记录最佳单次排名
func UpdateRecordBestSingleRank(dimension any) error {
	db := database.GetMySQL()

	// 开启事务
	tx := db.Begin()

	// 创建临时表
	err := tx.Exec("CREATE TEMPORARY TABLE temp_rank SELECT id FROM record_best_single WHERE dimension = ? ORDER BY record_duration", dimension).Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("创建临时表失败")
	}

	// 设置变量
	err = tx.Exec("SET @rank = 0").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("设置变量失败")
	}

	// 更新记录
	err = tx.Exec("UPDATE record_best_single AS r JOIN temp_rank AS tr ON r.id = tr.id SET r.rank = (@rank := @rank + 1)").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("更新记录失败")
	}

	// 删除临时表
	err = tx.Exec("DROP TEMPORARY TABLE IF EXISTS temp_rank").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("删除临时表失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}
