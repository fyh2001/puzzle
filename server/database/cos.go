package database

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"puzzle/config"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosClient struct {
	*cos.Client
}

var BaseURL = config.Settings.Cos.Url

func NewCosClient() *cos.Client {
	u, _ := url.Parse(BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Settings.Cos.SecretID,
			SecretKey: config.Settings.Cos.SecretKey,
		},
	})

	return client
}

func (client CosClient) Upload(file io.Reader, filePath string) error {
	_, err := client.Client.Object.Put(context.Background(), filePath, file, nil)
	if err != nil {
		return err
	}

	return nil
}

func (client *CosClient) Delete(filePath string) error {
	_, err := client.Client.Object.Delete(context.Background(), filePath)
	if err != nil {
		return err
	}

	return nil
}
