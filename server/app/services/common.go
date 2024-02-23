package services

import (
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
