package controllers

import (
	"net/http"
	"puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"github.com/gin-gonic/gin"
)

type UpdateLogController struct {
}

func (UpdateLogController) Insert(c *gin.Context) {
	var updateLog models.UpdateLog

	err := c.BindJSON(&updateLog)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail("[update_log]参数错误"))
		return
	}

	err = services.UpdateLog.Insert(&updateLog)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success("新增更新日志成功"))
}

func (UpdateLogController) List(c *gin.Context) {
	var updateLogReq models.UpdateLogReq

	err := c.Bind(&updateLogReq)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail("[update_log]参数错误"))
		return
	}

	updateLogListResp, err := services.UpdateLog.List(&updateLogReq)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success(updateLogListResp))
}
