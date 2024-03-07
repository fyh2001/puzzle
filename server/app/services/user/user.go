package user

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
	jwt "puzzle/utils/jwt"
	"strconv"
)

// UserRegister 用户注册
func UserRegister(u models.UserRegisterReq) error {

	// 判断用户名是否存在
	var tempUser models.User
	err := database.GetMySQL().Table("user").Where("username = ?", u.Username).First(&tempUser).Error
	if err == nil {
		return errors.New("用户名已存在")
	}

	// 判断昵称是否存在
	_ = database.GetMySQL().Table("user").Where("nickname = ?", u.Nickname).First(&tempUser).Error
	if tempUser.Id != 0 {
		return errors.New("昵称已存在")
	}

	snowflake := utils.Snowflake{}

	user := models.User{
		Id:       snowflake.NextVal(),
		Username: u.Username,
		Password: utils.MD5(u.Password),
		Nickname: u.Nickname,
		Status:   1,
	}

	return database.GetMySQL().Create(&user).Error
}

// UserLogin 用户登录
func UserLogin(u models.UserLoginReq) (models.UserLoginResp, error) {
	var loginResp models.UserLoginResp

	// 根据用户名获取用户信息
	var userInfo models.User
	err := database.GetMySQL().Table("user").Where("username = ?", u.Username).First(&userInfo).Error

	if err != nil {
		return loginResp, errors.New("用户不存在")
	}

	if userInfo.Password != utils.MD5(u.Password) {
		return loginResp, errors.New("密码错误")
	}

	if userInfo.Status == 2 {
		return loginResp, errors.New("用户已冻结")
	}

	if userInfo.Status == 3 {
		return loginResp, errors.New("用户已注销")
	}

	token, err := jwt.GenerateToken(userInfo.Id, userInfo.Username)
	if err != nil {
		return loginResp, errors.New("生成token失败")
	}

	// 返回用户信息
	loginResp.Token = token
	loginResp.User = models.UserResp{
		Id:        strconv.FormatInt(userInfo.Id, 10),
		Username:  userInfo.Username,
		Nickname:  userInfo.Nickname,
		Avatar:    userInfo.Avatar,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		Status:    userInfo.Status,
		CreatedAt: userInfo.CreatedAt,
		UpdatedAt: userInfo.UpdatedAt,
	}

	return loginResp, nil
}

// List 获取用户列表
func List(u models.UserReq) (models.UserListResp, error) {
	var userResp models.UserListResp

	if u.IdStr != "" {
		u.Id, _ = strconv.ParseInt(u.IdStr, 10, 64)
	}

	if u.OrderBy == "" {
		u.OrderBy = "id"
	}

	db := database.GetMySQL().Table("user").Order(u.OrderBy + " " + u.Sorted)

	if u.Id != 0 {
		db = db.Where("id = ?", u.Id)
	}

	if u.Username != "" {
		db = db.Where("username Like ?", "%"+u.Username+"%")
	}

	if u.Nickname != "" {
		db = db.Where("nickname Like ?", "%"+u.Nickname+"%")
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
		db = db.Where("created_at >= ? and created_at <= ?", u.DateRange[0], u.DateRange[1])
	}

	if len(u.Ids) > 0 {
		db = db.Where("id in (?)", u.Ids)
	}

	// 查询总数
	err := db.Count(&userResp.Total).Error
	if err != nil {
		return userResp, errors.New("查询失败")
	}

	// 分页
	if u.Pagination.Page > 0 && u.Pagination.PageSize > 0 {
		db = db.Scopes(utils.Paginate(&u.Pagination))
	}

	// 查询列表
	err = db.Find(&userResp.Records).Error
	if err != nil {
		return userResp, errors.New("查询失败")
	}

	return userResp, nil
}

// UpdateUser 更新用户
func Update(u models.User) error {

	u.Password = utils.MD5(u.Password)

	err := database.GetMySQL().Table("user").Updates(&u).Error
	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}
