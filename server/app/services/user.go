package services

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
	"strconv"

	"gorm.io/gorm"
)

type UserService interface {
	Register(u *models.UserRegisterReq) error
	Login(u *models.UserLoginReq) (models.UserLoginResp, error)
	List(u *models.UserReq) (models.UserListResp, error)
	Update(u *models.User) error
	applyFilters(db *gorm.DB, u *models.UserReq) *gorm.DB
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

func (UserServiceImpl) List(u *models.UserReq) (models.UserListResp, error) {
	var userResp models.UserListResp

	if u.IDStr != "" {
		u.ID, _ = strconv.ParseInt(u.IDStr, 10, 64)
	}

	if u.OrderBy == "" {
		u.OrderBy = "id"
	}

	db := database.GetMySQL().Model(&models.User{}).Order(u.OrderBy + " " + u.Sorted)

	db = applyFilters(db, u)

	if err := db.Count(&userResp.Total).Error; err != nil {
		return userResp, errors.New("查询失败")
	}

	if err := db.Find(&userResp.Records).Error; err != nil {
		return userResp, errors.New("查询失败")
	}

	return userResp, nil
}

func (UserServiceImpl) Update(u *models.User) error {
	if u.Password != "" {
		u.Password = utils.MD5(u.Password)
	}

	if u.Status == 3 {
		var user models.User

		err := database.GetMySQL().Model(&models.User{}).First(&user, user.ID).Error
		if err != nil {
			return errors.New("用户不存在")
		}

		u.Username = user.Username + "_del"
		u.Nickname = user.Nickname + "_del"
	}

	err := database.GetMySQL().Model(&models.User{}).Updates(u).Error
	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}

func applyFilters(db *gorm.DB, u *models.UserReq) *gorm.DB {
	if u.ID != 0 {
		db = db.Where("id = ?", u.ID)
	}

	if u.Username != "" {
		db = db.Where("username LIKE ?", "%"+u.Username+"%")
	}

	if u.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+u.Nickname+"%")
	}

	if u.AccoladeId != 0 {
		db = db.Where("accolade_id = ?", u.AccoladeId)
	}

	if u.Email != "" {
		db = db.Where("email = ?", u.Email)
	}

	if u.Phone != "" {
		db = db.Where("phone = ?", u.Phone)
	}

	if u.Status != 0 {
		db = db.Where("status = ?", u.Status)
	}

	if len(u.DateRange) == 2 && !u.DateRange[0].IsZero() && !u.DateRange[1].IsZero() {
		db = db.Where("created_at >= ? AND created_at <= ?", u.DateRange[0], u.DateRange[1])
	}

	if len(u.IDs) > 0 {
		db = db.Where("id IN (?)", u.IDs)
	}

	if u.Pagination.Page > 0 && u.Pagination.PageSize > 0 {
		db.Scopes(utils.Paginate(&u.Pagination))
	}

	return db
}
