package service

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Aws struct {
		S3 struct {
			BucketName string
			REGION     string
			//TODO AccessKeyID string
			//TODO SecretAccessKey string
		}
	}
}

func NewConfig() *Config {
	config := new(Config)
	config.Aws.S3.REGION = "ap-northeast-1"
	config.Aws.S3.BucketName = "genemes-pic"
	return config
}

type AwsS3Storage struct {
	Config *Config
	//MaxSize int32
	RetriesTimes int8
	UpLoader     *s3manager.Uploader
}

//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
func NewAwsS3Storage() AwsS3Storage {
	os.Setenv("AWS_PROFILE", test-account)
	config := NewConfig()
	newSession, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1"),
		CredentialsChainVerboseErrors: aws.Bool(true),
		HTTPClient:                    &http.Client{Timeout: 10 * time.Second},
		Credentials:                   credentials.NewSharedCredentials("", "test-account"),
	})
	if err != nil {
		println("S3 Storage Session init failed")
	}
	uploader := s3manager.NewUploader(newSession)
	return AwsS3Storage{Config: config, RetriesTimes: 2, UpLoader: uploader}
}

func (storage *AwsS3Storage) UploadMemePic(file multipart.File, fileName string, extension string) (url string, err error) {
	if fileName == "" {
		return "", errors.New("fileName is required")
	}
	var contentType string

	switch extension {
	case "jpg":
		contentType = "image/jpeg"
	case "jpeg":
		contentType = "image/jpeg"
	case "gif":
		contentType = "image/gif"
	case "png":
		contentType = "image/png"
	default:
		return "", errors.New("this extension is invalid")
	}
	result, err := storage.UpLoader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        file,
		Bucket:      aws.String(storage.Config.Aws.S3.BucketName),
		Key:         aws.String("niconicocsc"),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file,%v", err)
	}
	return result.Location, nil
}
