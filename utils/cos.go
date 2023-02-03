package utils

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"tiktok/config"
	"tiktok/log"
	"time"
)

var cosClient *cos.Client
var ak string
var sk string

func InitCos() {
	u, err := url.Parse(config.Conf.GetString("bucket.url"))
	if err != nil {
		log.Fatalln(err)
	}
	cosClient = cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Conf.GetString("bucket.secret.id"),
			SecretKey: config.Conf.GetString("bucket.secret.key"),
		},
	})
	ak = config.Conf.GetString("bucket.secret.id")
	sk = config.Conf.GetString("bucket.secret.key")
}

func IsExist(ctx context.Context, path string) bool {
	ok, err := cosClient.Object.IsExist(ctx, path)
	if err != nil {
		log.Errorf("fail to check if object exist:%v", err)
		return true
	}
	return ok
}

func Upload(ctx context.Context, ext string, data io.Reader) (string, error) {
	// find a unique name
	var uu uuid.UUID
	var filename string
	for {
		uu = uuid.New()
		filename = fmt.Sprintf("upload/%s%s", uu.String(), ext)
		if !IsExist(ctx, filename) {
			break
		}
	}

	// upload
	_, err := cosClient.Object.Put(ctx, filename, data, nil)
	if err != nil {
		return "", err
	}

	return uu.String(), nil
}

func GenUrl(ctx context.Context, filename string) string {
	u, err := cosClient.Object.GetPresignedURL(ctx, http.MethodGet, filename, ak, sk, time.Hour*24, nil)
	if err != nil {
		log.Errorln(err)
		return ""
	}
	return u.String()
}
