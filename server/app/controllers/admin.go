package controllers

import (
	"fmt"
	"puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (AdminController) ListUserData(c *gin.Context) {
	var userReq models.UserReq

	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	userReq.Id, _ = strconv.ParseInt(userReq.IdStr, 10, 64)

	userListResp, err := services.User.List(&userReq)

	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(userListResp))
}

func (AdminController) UpdateUserData(c *gin.Context) {
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

	err = services.User.Update(&user)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success("更新成功"))
}

func (AdminController) ListRecordData(c *gin.Context) {
	var recordReq models.RecordReq
	err := c.ShouldBindJSON(&recordReq)
	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordListResp, err := services.Record.List(&recordReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordListResp))
}

func (AdminController) UpdateRecordData(c *gin.Context) {
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

	err = services.Record.Update(&record)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success("更新成功"))
}

func (AdminController) ListRecordBestSingleData(c *gin.Context) {
	var recordBestSingleReq models.RecordBestSingleReq
	err := c.ShouldBindJSON(&recordBestSingleReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordBestSingleListResp, err := services.RecordBestSingle.List(&recordBestSingleReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordBestSingleListResp))
}

func (AdminController) ListRecordBestAverageData(c *gin.Context) {
	var recordBestAverageReq models.RecordBestAverageReq
	err := c.ShouldBindJSON(&recordBestAverageReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordBestAverageListResp, err := services.RecordBestAverage.List(&recordBestAverageReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordBestAverageListResp))
}

func (AdminController) ListRecordBestStepData(c *gin.Context) {
	var recordBestStepReq models.RecordBestStepReq
	err := c.ShouldBindJSON(&recordBestStepReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	recordBestStepListResp, err := services.RecordBestStep.List(&recordBestStepReq)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(recordBestStepListResp))
}
