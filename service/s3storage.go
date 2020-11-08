package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kataras/iris/v12"
)

type S3Storage struct {
	BucketName string
	REGION     string
	//MaxSize int32
	RetriesTimes int8
	UpLoader     *s3manager.Uploader
}

func (storage *S3Storage) New() S3Storage {
	newSession, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")})
	if err != nil {
		println("S3 Storage Session init failed")
	}
	uploader := s3manager.NewUploader(newSession)
	return S3Storage{BucketName: "genemes-pic", REGION: "ap-northeast-1", RetriesTimes: 2, UpLoader: uploader}
}

func (storage *S3Storage) StoreMemePic(ctx iris.Context) {

}
