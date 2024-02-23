package controllers

import (
	"puzzle/app/common/result"
	adminAuthorizationService "puzzle/app/services/admin-authorization"

	"github.com/gin-gonic/gin"
)

type AdminAuthorization struct {
	Code string `json:"code"` // 验证码
}

func ResetOtp(c *gin.Context) {
	url, err := adminAuthorizationService.GenerateSecretKey()
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(url))
}

func Authorization(c *gin.Context) {
	var adminAuthorization AdminAuthorization

	// 获取参数
	if err := c.ShouldBindJSON(&adminAuthorization); err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	// 校验验证码
	adminAuthorizationResp, err := adminAuthorizationService.Authorization(adminAuthorization.Code)

	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(adminAuthorizationResp))
}

func GetUrl(c *gin.Context) {
	url, err := adminAuthorizationService.GetUrl()
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(url))
}
