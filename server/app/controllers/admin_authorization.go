package controllers

import (
	"puzzle/app/common/result"
	"puzzle/app/models"
	"puzzle/app/services"

	"github.com/gin-gonic/gin"
)

type AdminAuthorizationController struct{}

type AdminAuthorizationResp struct {
	Code string `json:"code"` // 验证码
}

func (AdminAuthorizationController) ResetOtp(c *gin.Context) {
	url, err := services.AdminAuthorization.GenerateSecretKey()
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	secretResp := models.SecretResp{
		SecretUrl: url,
	}

	c.JSON(200, result.Success(secretResp))
}

func (AdminAuthorizationController) Authorization(c *gin.Context) {
	var adminAuthorization AdminAuthorizationResp

	// 获取参数
	if err := c.ShouldBindJSON(&adminAuthorization); err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	// 校验验证码
	adminAuthorizationResp, err := services.AdminAuthorization.Authorization(adminAuthorization.Code)

	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	c.JSON(200, result.Success(adminAuthorizationResp))
}

func (AdminAuthorizationController) GetUrl(c *gin.Context) {
	url, err := services.AdminAuthorization.GetUrl()
	if err != nil {
		c.JSON(200, result.Fail(err.Error()))
		return
	}

	secretResp := models.SecretResp{
		SecretUrl: url,
	}

	c.JSON(200, result.Success(secretResp))
}
