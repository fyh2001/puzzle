package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	userService "puzzle/app/services/user"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerReq models.UserRegisterReq

	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = userService.UserRegister(registerReq)
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
	userInfo, err := userService.UserLogin(loginReq)
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
	userList, err := userService.List(userReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(userList))
}
