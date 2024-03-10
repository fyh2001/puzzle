package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	userService "puzzle/app/services/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerReq models.UserRegisterReq

	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = userService.UserRegister(&registerReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("注册成功"))
}

func Login(c *gin.Context) {
	var loginReq models.UserLoginReq

	// 参数校验
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	// 登录
	userInfo, err := userService.UserLogin(&loginReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(userInfo))
}

func ListUser(c *gin.Context) {
	var userReq models.UserReq

	// 参数校验
	err := c.ShouldBind(&userReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	// 查询用户列表
	userList, err := userService.List(&userReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(userList))
}

func GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")

	userInfo, err := userService.GetUserInfo(userId.(int64))
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(userInfo))
}

func UpdateAvatar(c *gin.Context) {
	var user models.User

	userId, _ := c.Get("userId")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	user.Id = userId.(int64)

	// 更新头像
	err = userService.UpdateAvatar(&user, avatar)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("更新成功"))
}

func UpdateUser(c *gin.Context) {
	var userReq models.UserReq

	err := c.ShouldBind(&userReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	id, _ := strconv.ParseInt(userReq.IdStr, 10, 64)

	user := models.User{
		Id:         id,
		Username:   userReq.Username,
		Nickname:   userReq.Nickname,
		AccoladeId: userReq.AccoladeId,
		Email:      userReq.Email,
		Phone:      userReq.Phone,
		Password:   userReq.Password,
		Status:     userReq.Status,
	}

	// 更新用户
	err = userService.Update(&user)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("更新成功"))
}
