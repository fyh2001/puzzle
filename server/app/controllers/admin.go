package controllers

import (
	"fmt"
	"puzzle/app/common/result"
	"puzzle/app/models"

	recordService "puzzle/app/services/record"
	recordBestAverageSerivce "puzzle/app/services/record-best-average"
	recordBestSingleSerivce "puzzle/app/services/record-best-single"
	recordBestStepSerivce "puzzle/app/services/record-best-step"
	userService "puzzle/app/services/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListUserData(c *gin.Context) {
	var userReq models.UserReq

	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	userReq.Id, _ = strconv.ParseInt(userReq.IdStr, 10, 64)

	userListResp, err := userService.List(&userReq)

	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(userListResp))
}

func UpdateUserData(c *gin.Context) {
	var userReq models.UserReq
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	id, _ := strconv.ParseInt(userReq.IdStr, 10, 64)

	user := models.User{
		Id:       id,
		Username: userReq.Username,
		Password: userReq.Password,
		Nickname: userReq.Nickname,
		Phone:    userReq.Phone,
		Email:    userReq.Email,
		Status:   userReq.Status,
	}

	err = userService.Update(&user)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success("更新成功"))
}

func ListRecordData(c *gin.Context) {
	var recordReq models.RecordReq
	err := c.ShouldBindJSON(&recordReq)
	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordListResp, err := recordService.List(&recordReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordListResp))
}

func UpdateRecordData(c *gin.Context) {
	var recordReq models.RecordReq
	err := c.ShouldBindJSON(&recordReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	id, _ := strconv.ParseInt(recordReq.IdStr, 10, 64)
	userId, _ := strconv.ParseInt(recordReq.UserIdStr, 10, 64)
	idx, _ := strconv.ParseInt(recordReq.IdxStr, 10, 64)

	record := models.Record{
		Id:        id,
		UserId:    userId,
		Dimension: recordReq.Dimension,
		Type:      recordReq.Type,
		Duration:  recordReq.Duration,
		Step:      recordReq.Step,
		Scramble:  recordReq.Scramble,
		Solution:  recordReq.Solution,
		Idx:       idx,
		Status:    recordReq.Status,
	}

	err = recordService.Update(&record)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success("更新成功"))
}

func ListRecordBestSingleData(c *gin.Context) {
	var recordBestSingleReq models.RecordBestSingleReq
	err := c.ShouldBindJSON(&recordBestSingleReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordBestSingleListResp, err := recordBestSingleSerivce.List(&recordBestSingleReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordBestSingleListResp))
}

func ListRecordBestAverageData(c *gin.Context) {
	var recordBestAverageReq models.RecordBestAverageReq
	err := c.ShouldBindJSON(&recordBestAverageReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordBestAverageListResp, err := recordBestAverageSerivce.List(&recordBestAverageReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordBestAverageListResp))
}

func ListRecordBestStepData(c *gin.Context) {
	var recordBestStepReq models.RecordBestStepReq
	err := c.ShouldBindJSON(&recordBestStepReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordBestStepListResp, err := recordBestStepSerivce.List(&recordBestStepReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordBestStepListResp))
}
