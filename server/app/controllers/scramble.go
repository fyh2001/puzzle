package controllers

import (
	"puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"github.com/gin-gonic/gin"
)

type ScrambleController struct{}

func (ScrambleController) GetNewScamble(c *gin.Context) {
	var req models.GetNewScambleReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	userId, _ := c.Get("userId")
	req.UserId = userId.(int64)

	scramble, err := services.Scramble.GetNewScamble(&req)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(scramble))
}

func (ScrambleController) GetUserScramble(c *gin.Context) {
	var req models.GetNewScambleReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	userId, _ := c.Get("userId")
	req.UserId = userId.(int64)

	scramble, err := services.Scramble.GetUserScramble(&req)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(scramble))
}
