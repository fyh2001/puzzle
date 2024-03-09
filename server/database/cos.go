package database

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosClient struct {
	*cos.Client
}

var BaseURL = "https://defo1215-1307771416.cos.ap-guangzhou.myqcloud.com"

func NewCosClient() *cos.Client {
	u, _ := url.Parse(BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDCE8nEQN6bwQhkZ9VcU9amjRIvb6QZWGn",
			SecretKey: "PcfVBe1B8uVu66vYyCE6lFXeYFk294nJ",
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
