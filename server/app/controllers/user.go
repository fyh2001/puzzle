package controllers

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"

	"gorm.io/gorm"
)

type UserService interface {
	Register(u *models.UserRegisterReq) error
}

type UserServiceImpl struct{}

const (
	StatusNormal int8 = 1 + iota
	StatusFrozen
	StatusDeleted
)

func (UserServiceImpl) Register(u *models.UserRegisterReq) error {

	var tempUser models.User

	// 检查用户名重复
	row := database.GetMySQL().Model(&models.User{}).Where("username = ? AND status != ?", u.Username, StatusDeleted).First(&tempUser).RowsAffected
	if row > 0 {
		return errors.New("用户名已存在")
	}

	// 检查昵称重复
	row = database.GetMySQL().Model(&models.User{}).Where("nickname = ? AND status != ?", u.Nickname, StatusDeleted).First(&tempUser).RowsAffected
	if row > 0 {
		return errors.New("昵称已存在")
	}

	user := &models.User{
		ID:       (&utils.Snowflake{}).NextVal(),
		Username: u.Username,
		Password: utils.MD5(u.Password),
		Nickname: u.Nickname,
		Status:   StatusNormal,
	}

	return database.GetMySQL().Create(user).Error
}

func (UserServiceImpl) Login(u *models.UserLoginReq) (models.UserLoginResp, error) {

	var loginResp models.UserLoginResp

	var userInfo models.User
	err := database.GetMySQL().Model(&models.User{}).Where("username = ?", u.Username).First(&userInfo).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return loginResp, errors.New("用户不存在")
	}

	if userInfo.Password != utils.MD5(u.Password) {
		return loginResp, errors.New("用户名或密码错误")
	}

	if userInfo.Status == StatusFrozen {
		return loginResp, errors.New("用户已冻结")
	}

	if userInfo.Status == StatusDeleted {
		return loginResp, errors.New("用户已注销")
	}

	return loginResp, nil
}
