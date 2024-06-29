package controllers

import (
	"puzzle/app/common"
	"puzzle/app/models"
	"puzzle/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) Register(c *gin.Context) {

	var registerReq models.UserRegisterReq

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	if err := services.User.Register(&registerReq); err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success("注册成功"))
}

func (UserController) Login(c *gin.Context) {

	var loginReq models.UserLoginReq

	// 参数校验
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	// 登录
	userInfo, err := services.User.Login(&loginReq)
	if err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success(userInfo))
}

func (UserController) List(c *gin.Context) {

	var userReq models.UserReq

	// 参数校验
	if err := c.ShouldBind(&userReq); err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	// 查询用户列表
	userList, err := services.User.List(&userReq)
	if err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success(userList))
}

func (UserController) Update(c *gin.Context) {
	var userReq models.UserReq

	err := c.ShouldBind(&userReq)
	if err != nil {
		c.JSON(200, common.HttpResult.Fail("参数错误"))
		return
	}

	id, _ := strconv.ParseInt(userReq.IDStr, 10, 64)

	user := &models.User{
		ID:         id,
		Username:   userReq.Username,
		Nickname:   userReq.Nickname,
		AccoladeId: userReq.AccoladeId,
		Email:      userReq.Email,
		Phone:      userReq.Phone,
		Password:   userReq.Password,
		Status:     userReq.Status,
	}

	// 更新用户
	err = services.User.Update(user)
	if err != nil {
		c.JSON(200, common.HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, common.HttpResult.Success("更新成功"))
}
