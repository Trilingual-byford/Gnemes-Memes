package test

import (
	"fmt"
	"github.com/corona10/goimagehash"
	"github.com/smartystreets/assertions"
	"image/png"
	"myapp/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestExtensionUtil(t *testing.T) {
	ext := filepath.Ext("test.png")
	fmt.Print(ext)
	assertions.ShouldEqual(".png", ext)
}

func TestPicUtil(t *testing.T) {
	file, _ := os.Open("/Users/byford/go/src/Gnemes-Memes/asset/testPic.png")
	//defer file.Close()
	//img, err := jpeg.Decode(file)
	img, err := png.Decode(file)
	hash, err2 := goimagehash.AverageHash(img)
	fmt.Printf("Pic hash:%v\n", hash)
	fmt.Printf("Pic hash error :%v\n", err)
	fmt.Printf("Pic hash error2 :%v\n", err2)
	differenceHash, err3 := goimagehash.DifferenceHash(img)
	fmt.Printf(" Pic DifferenceHash:%v", differenceHash)
	fmt.Printf("Pic differenceHash error2 :%v", err3)

	assertions.ShouldNotBeNil(hash)
}
func TestPicUtilFuncTest(t *testing.T) {
	file, _ := os.Open("/Users/byford/go/src/Gnemes-Memes/asset/testPic.png")
	hash, err := utils.GetPicHash(file, ".png")
	assertions.ShouldBeNil(err)
	assertions.ShouldNotBeNil(hash)
	jpegFile, _ := os.Open("/Users/byford/go/src/Gnemes-Memes/asset/pic/funny-wholesome-animal-memes-1.jpeg")
	jpegHash, err2 := utils.GetPicHash(jpegFile, ".jpeg")
	assertions.ShouldBeNil(err2)
	assertions.ShouldNotBeNil(jpegHash)
	jpgFile, _ := os.Open("/Users/byford/go/src/Gnemes-Memes/asset/pic/funny-wholesome-animal-memes-2.jpg")
	jpgHash, err3 := utils.GetPicHash(jpgFile, ".jpg")
	assertions.ShouldBeNil(err3)
	assertions.ShouldNotBeNil(jpgHash)

}
