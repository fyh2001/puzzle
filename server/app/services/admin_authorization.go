package services

import (
	"context"
	"errors"
	"fmt"
	"puzzle/app/models"
	"puzzle/database"
	jwt "puzzle/utils/jwt"

	"github.com/pquerna/otp/totp"
)

type AdminAuthorizationService interface {
	GenerateSecretKey() (string, error)
	Authorization(code string) (models.AdminAuthorizationResp, error)
	GetUrl() (string, error)
}

type AdminAuthorizationImpl struct{}

var option = totp.GenerateOpts{
	Issuer:      "puzzle",
	AccountName: "defo1215.cn",
}

// GenerateSecretKey 生成一个新的密钥
func (AdminAuthorizationImpl) GenerateSecretKey() (string, error) {
	key, err := totp.Generate(option)
	if err != nil {
		return "", errors.New("生成密钥失败")
	}

	url := key.URL()

	// 将url存入redis
	err = database.GetRedis().Set(context.Background(), "admin:opt_url", url, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("opt_url存入redis失败")
	}

	// 将secret存入redis
	err = database.GetRedis().Set(context.Background(), "admin:opt_secret", key.Secret(), 0).Err()
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("opt_secret存入redis失败")
	}

	return url, nil
}

// Authorization 验证验证码并获取token
func (AdminAuthorizationImpl) Authorization(code string) (models.AdminAuthorizationResp, error) {
	var adminAuthorizationResp models.AdminAuthorizationResp

	// 从redis获取secret
	secret := database.GetRedis().Get(context.Background(), "admin:opt_secret").Val()

	// 从redis获取url
	url := database.GetRedis().Get(context.Background(), "admin:opt_url").Val()

	// 验证验证码
	ok := totp.Validate(code, secret)
	if !ok {
		return models.AdminAuthorizationResp{}, errors.New("验证码错误")
	}

	// 获取token
	token, err := jwt.GenerateAdminToken(models.AdminAuthorization{OptUrl: url, Secret: secret})
	if err != nil {
		return models.AdminAuthorizationResp{}, errors.New("生成token失败")
	}

	adminAuthorizationResp.OptUrl = url
	adminAuthorizationResp.Token = token

	return adminAuthorizationResp, nil
}

// GetUrl 获取url
func (AdminAuthorizationImpl) GetUrl() (string, error) {
	url, err := database.GetRedis().Get(context.Background(), "admin:opt_url").Result()
	if err != nil {
		return "", errors.New("获取opt_url失败")
	}

	return url, nil
}
