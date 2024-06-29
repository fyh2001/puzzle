package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"github.com/gin-gonic/gin"
)

type RecordBestAverageController struct{}

func (RecordBestAverageController) List(c *gin.Context) {
	var recordReq models.RecordBestAverageReq
	err := c.ShouldBind(&recordReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	recordList, err := services.RecordBestAverage.List(&recordReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(recordList))
}
