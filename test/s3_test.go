package test

import (
	"fmt"
	"github.com/smartystreets/assertions"
	"myapp/service"
	"os"
	"testing"
)

func TestS3(t *testing.T) {
	file, err := os.Open("/Users/byford/go/src/Gnemes-Memes/asset/testPic.png")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer file.Close()
	awsS3 := service.NewAwsS3Storage()
	picUrl, err := awsS3.UploadMemePic(file, "test", "png")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(picUrl)
	assertions.ShouldNotBeIn(picUrl)
}
