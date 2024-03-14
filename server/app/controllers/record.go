package controllers

import (
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"puzzle/utils"

	"github.com/gin-gonic/gin"
)

type RecordController struct{}

func (RecordController) Insert(c *gin.Context) {
	var record models.Record
	err := c.ShouldBind(&record)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	encryptionParams := utils.EncryptionParams{
		Dimension: record.Dimension,
		RandomIdx: record.Idx,
		StepCount: record.Step,
		Scramble:  record.Scramble,
		Solution:  record.Solution,
	}

	if !encryptionParams.VerifyScramble() {
		c.JSON(200, HttpResult.Fail("参数错误!"))
		return
	}

	// 获取用户ID
	userId, _ := c.Get("userId")
	record.UserId = userId.(int64)

	err = services.Record.Insert(&record)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("添加成功"))
}

func (RecordController) List(c *gin.Context) {
	var recordReq models.RecordReq
	err := c.ShouldBind(&recordReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	recordList, err := services.Record.List(&recordReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(recordList))
}

func (RecordController) Update(c *gin.Context) {
	var record models.Record
	err := c.ShouldBind(&record)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = services.Record.Update(&record)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("修改成功"))
}
