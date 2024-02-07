package controllers

import (
	"fmt"
	HttpResult "puzzle/app/common/result"
	"puzzle/app/models"
	recordService "puzzle/app/services/record"
	"puzzle/utils"

	"github.com/gin-gonic/gin"
)

func InsertRecord(c *gin.Context) {
	var record models.Record
	err := c.ShouldBind(&record)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	fmt.Printf("record: %v\n", record)

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

	err = recordService.Insert(record)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("添加成功"))
}

func ListRecord(c *gin.Context) {
	var recordReq models.RecordReq
	err := c.ShouldBind(&recordReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	recordList, err := recordService.List(recordReq)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success(recordList))
}

func UpdateRecord(c *gin.Context) {
	var record models.Record
	err := c.ShouldBind(&record)
	if err != nil {
		c.JSON(200, HttpResult.Fail("参数错误"))
		return
	}

	err = recordService.Update(record)
	if err != nil {
		c.JSON(200, HttpResult.Fail(err.Error()))
		return
	}

	c.JSON(200, HttpResult.Success("修改成功"))
}
