package services

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"puzzle/app/models"
	"puzzle/app/services/options"
	"puzzle/app/types"
	"puzzle/database"
	"puzzle/utils"
	"strconv"
)

type UserService struct{}

func (UserService) Register(req *models.UserRegisterReq) error {

	var tempUser models.User

	// 检查用户名重复
	row := database.GetMySQL().Model(&models.User{}).Where("username = ?", req.Username).First(&tempUser).RowsAffected
	if row > 0 {
		return errors.New("用户名已存在")
	}

	// 检查昵称重复
	row = database.GetMySQL().Model(&models.User{}).Where("nickname = ?", req.Nickname).First(&tempUser).RowsAffected
	if row > 0 {
		return errors.New("昵称已存在")
	}

	user := &models.User{
		ID:       (&utils.Snowflake{}).NextVal(),
		Username: req.Username,
		Password: utils.MD5(req.Password),
		Nickname: req.Nickname,
		Status:   types.UserStatusNormal,
	}

	return database.GetMySQL().Create(user).Error
}

func (UserService) Login(req *models.UserLoginReq) (resp *models.UserLoginResp, err error) {

	var userInfo models.User
	err = database.GetMySQL().Model(&models.User{}).Where("username = ?", req.Username).First(&userInfo).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &models.UserLoginResp{}, errors.New("用户不存在")
	}

	if userInfo.Password != utils.MD5(req.Password) {
		return resp, errors.New("用户名或密码错误")
	}

	if userInfo.Status == types.UserStatusFrozen {
		return &models.UserLoginResp{}, errors.New("用户已冻结")
	}

	return &models.UserLoginResp{}, nil
}

func (UserService) List(req *models.UserReq, reqUid int64) (resp *models.UserListResp, err error) {

	if req.IDStr != "" {
		req.ID, _ = strconv.ParseInt(req.IDStr, 10, 64)
	}

	if reqUid != types.ReqAdminUid {
		// 普通用户
		if req.ID != 0 && (req.ID != reqUid || len(req.IDs) != 0) {
			return &models.UserListResp{}, errors.New("权限不足")
		}

		req.ID = reqUid
	}

	db := database.GetMySQL().Model(&models.User{})

	filters := []options.QueryOption{
		options.WithQueryID64(req.ID),
		options.WithQueryIDs64(req.IDs),
		options.WithQueryAccoladeID(req.AccoladeId),
		options.WithQueryUsername(req.Username),
		options.WithQueryNickname(req.Nickname),
		options.WithQueryEmail(req.Email),
		options.WithQueryPhone(req.Phone),
		options.WithQueryStatus(int(req.Status)),
		options.WithQueryDateRange(req.DateRange),
		options.WithQueryPagination(req.Pagination),
		options.WithQueryOrderBy(req.OrderBy, req.Sorted),
	}

	options.ApplyQueryFilters(db, filters...)

	if err := db.Count(&resp.Total).Error; err != nil {
		return &models.UserListResp{}, errors.New("查询失败")
	}

	if err := db.Preload(clause.Associations).Find(&resp.Records).Error; err != nil {
		return &models.UserListResp{}, errors.New("查询失败")
	}

	return resp, nil
}

func (UserService) Update(req []*models.User, reqUid int64) error {

	//if u.Status == 3 {
	//	var user models.User
	//
	//	err := database.GetMySQL().Model(&models.User{}).First(&user, user.ID).Error
	//	if err != nil {
	//		return errors.New("用户不存在")
	//	}
	//
	//	u.Username = user.Username + "_del"
	//	u.Nickname = user.Nickname + "_del"
	//}

	if reqUid != types.ReqAdminUid {
		if len(req) != 0 || (req[0].ID != 0 && req[0].ID != reqUid) {
			return errors.New("权限不足")
		}
	}

	for i, _ := range req {
		req[i].Password = utils.MD5(req[i].Password)
	}

	if err := database.GetMySQL().Model(&models.User{}).Updates(req).Error; err != nil {
		return errors.New("更新失败")
	}

	return nil
}

func (UserService) Delete(req []*models.User, reqUid int64) error {

	if reqUid != types.ReqAdminUid {
		if len(req) != 0 || (req[0].ID != 0 && req[0].ID != reqUid) {
			return errors.New("权限不足")
		}
	}

	if err := database.GetMySQL().Model(&models.User{}).Delete(req).Error; err != nil {
		return errors.New("注销失败")
	}

	return nil
}
