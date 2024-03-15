package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"github.com/gin-gonic/gin"
)

type NotificationController struct{}

func (NotificationController) Insert(c *gin.Context) {

	var notification models.NotificationReq

	err := c.ShouldBind(&notification)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = services.Notification.Insert(&notification)

	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("发布成功"))
}

func (NotificationController) List(c *gin.Context) {

	var notificationReq models.NotificationReq

	err := c.ShouldBind(&notificationReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	notificationListResp, err := services.Notification.List(&notificationReq)

	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(notificationListResp))
}

func (NotificationController) Update(c *gin.Context) {

	var notification models.NotificationReq

	err := c.ShouldBind(&notification)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = services.Notification.Update(&notification)

	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("修改成功"))
}

func (NotificationController) ReadAll(c *gin.Context) {

	var notificationReq models.NotificationReq

	err := c.ShouldBind(&notificationReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = services.Notification.ReadAll(&notificationReq)

	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("一键已读成功"))
}
