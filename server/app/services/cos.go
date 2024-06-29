package services

import (
	"errors"
	"mime/multipart"
	"puzzle/database"
	"strings"

	"github.com/google/uuid"
)

type CosService interface {
	UploadAvatar(file *multipart.FileHeader) (string, error)
	DeleteAvatar(fullFilePath string) error
}

type CosImpl struct{}

var baseUrl = map[string]string{
	"avatar": "/avatar/",
}

func (CosImpl) UploadAvatar(file *multipart.FileHeader) (string, error) {
	var filePath string

	if file.Size > 5*1024*1024 {
		return "", errors.New("图片大小不能超过5Mb")
	}

	avatar, err := file.Open()
	if err != nil {
		return "", errors.New("图片打开失败")
	}
	defer avatar.Close()

	filePath = baseUrl["avatar"] + uuid.NewString() + "." + strings.Split(file.Filename, `.`)[1]

	client := database.CosClient{
		Client: database.NewCosClient(),
	}

	err = client.Upload(avatar, filePath)
	if err != nil {
		return "", errors.New("图片上传失败")
	}

	fullFilePath := database.BaseURL + filePath

	return fullFilePath, nil
}

func (CosImpl) DeleteAvatar(fullFilePath string) error {
	filePath := strings.Split(fullFilePath, database.BaseURL)[1]

	client := database.CosClient{
		Client: database.NewCosClient(),
	}

	err := client.Delete(filePath)
	if err != nil {
		return errors.New("图片删除失败")
	}

	return nil
}
