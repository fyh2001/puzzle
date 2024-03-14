package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"github.com/gin-gonic/gin"
)

type NotificationUserStatusController struct{}

func (NotificationUserStatusController) Insert(c *gin.Context) {

	var notificationUserStatus models.NotificationUserStatusReq

	err := c.ShouldBind(&notificationUserStatus)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	userId, _ := c.Get("userId")
	notificationUserStatus.UserId = userId.(int64)

	err = services.NotificationUserStatus.Insert(&notificationUserStatus)

	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("已读成功"))
}
