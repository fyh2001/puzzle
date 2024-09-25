package controllers

import (
	"puzzle/app/common"
	"puzzle/app/models"
	"puzzle/app/services"
	"puzzle/app/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) Register(c *gin.Context) {

	var req models.UserRegisterReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	if err := services.User.Register(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success("注册成功"))
}

func (UserController) Login(c *gin.Context) {

	var req models.UserLoginReq

	// 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	// 登录
	userInfo, err := services.User.Login(&req)
	if err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success(userInfo))
}

func (UserController) List(c *gin.Context) {

	var req models.UserReq

	// 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	reqUid, _ := c.Get("userID")

	// 查询用户列表
	userList, err := services.User.List(&req, reqUid.(int64))
	if err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success(userList))
}

func (UserController) Update(c *gin.Context) {

	var req []models.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	reqUid, _ := c.Get("userID")

	var reqSlice []*models.User

	for _, r := range req {
		id, _ := strconv.ParseInt(r.IDStr, 10, 64)

		user := &models.User{
			ID:         id,
			Username:   r.Username,
			Nickname:   r.Nickname,
			AccoladeId: r.AccoladeId,
			Email:      r.Email,
			Phone:      r.Phone,
			Password:   r.Password,
			Status:     r.Status,
		}

		reqSlice = append(reqSlice, user)
	}

	// 更新用户
	if err := services.User.Update(reqSlice, reqUid.(int64)); err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success("更新成功"))
}

func (UserController) Delete(c *gin.Context) {

	var req []models.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	reqUid, _ := c.Get("userID")

	var reqSlice []*models.User

	for _, r := range req {
		id, _ := strconv.ParseInt(r.IDStr, 10, 64)

		user := &models.User{
			ID:         id,
			Username:   r.Username,
			Nickname:   r.Nickname,
			AccoladeId: r.AccoladeId,
			Email:      r.Email,
			Phone:      r.Phone,
			Password:   r.Password,
			Status:     r.Status,
		}

		reqSlice = append(reqSlice, user)
	}

	// 注销用户
	if err := services.User.Delete(reqSlice, reqUid.(int64)); err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success("注销成功"))

}

func (UserController) ListByAdmin(c *gin.Context) {

	var req models.UserReq

	// 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	// 查询用户列表
	userList, err := services.User.List(&req, types.ReqAdminUid)
	if err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success(userList))
}

func (UserController) UpdateByAdmin(c *gin.Context) {

	var req []models.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	var reqSlice []*models.User

	for _, r := range req {
		id, _ := strconv.ParseInt(r.IDStr, 10, 64)

		user := &models.User{
			ID:         id,
			Username:   r.Username,
			Nickname:   r.Nickname,
			AccoladeId: r.AccoladeId,
			Email:      r.Email,
			Phone:      r.Phone,
			Password:   r.Password,
			Status:     r.Status,
		}

		reqSlice = append(reqSlice, user)
	}

	// 更新用户
	if err := services.User.Update(reqSlice, types.ReqAdminUid); err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success("更新成功"))
}

func (UserController) DeleteByAdmin(c *gin.Context) {

	var req []models.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	var reqSlice []*models.User

	for _, r := range req {
		id, _ := strconv.ParseInt(r.IDStr, 10, 64)

		user := &models.User{
			ID:         id,
			Username:   r.Username,
			Nickname:   r.Nickname,
			AccoladeId: r.AccoladeId,
			Email:      r.Email,
			Phone:      r.Phone,
			Password:   r.Password,
			Status:     r.Status,
		}

		reqSlice = append(reqSlice, user)
	}

	// 注销用户
	if err := services.User.Delete(reqSlice, types.ReqAdminUid); err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success("注销成功"))
}
