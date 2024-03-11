package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	notificationService "puzzle/app/services/notification"

	"github.com/gin-gonic/gin"
)

func InsertNotification(c *gin.Context) {

	var notification models.NotificationReq

	err := c.ShouldBind(&notification)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = notificationService.Insert(&notification)

	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("发布成功"))
}

func ListNotification(c *gin.Context) {

	var notificationReq models.NotificationReq

	err := c.ShouldBind(&notificationReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	notificationListResp, err := notificationService.List(&notificationReq)

	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(notificationListResp))
}
