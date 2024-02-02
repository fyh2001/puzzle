package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	recordService "puzzle/app/services/record"

	"github.com/gin-gonic/gin"
)

func InsertRecord(c *gin.Context) {
	var record models.Record
	err := c.ShouldBind(&record)
	if err != nil {
		HttpResult.Fail("参数错误")
		return
	}

	err = recordService.Insert(record)
	if err != nil {
		HttpResult.Fail(err.Error())
		return
	}

	HttpResult.Success("新增成功")
}

func ListRecord(c *gin.Context) {
	var recordReq models.RecordReq
	err := c.ShouldBind(&recordReq)
	if err != nil {
		HttpResult.Fail("参数错误")
		return
	}

	recordList, err := recordService.List(recordReq)
	if err != nil {
		HttpResult.Fail(err.Error())
		return
	}

	HttpResult.Success(recordList)
}

func UpdateRecord(c *gin.Context) {
	var record models.Record
	err := c.ShouldBind(&record)
	if err != nil {
		HttpResult.Fail("参数错误")
		return
	}

	err = recordService.Update(record)
	if err != nil {
		HttpResult.Fail(err.Error())
		return
	}

	HttpResult.Success("更新成功")
}