package controllers

import (
	"puzzle/app/common/result"
	"puzzle/app/models"
	scrambleService "puzzle/app/services/scramble"

	"github.com/gin-gonic/gin"
)

func GetNewScamble(c *gin.Context) {
	var req models.GetNewScambleReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	userId, _ := c.Get("userId")
	req.UserId = userId.(int64)

	scramble, err := scrambleService.GetNewScamble(req)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(scramble))
}

func GetUserScramble(c *gin.Context) {
	var req models.GetNewScambleReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	userId, _ := c.Get("userId")
	req.UserId = userId.(int64)

	scramble, err := scrambleService.GetUserScramble(req)
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(scramble))
}
