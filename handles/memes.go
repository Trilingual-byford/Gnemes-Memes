package handles

import (
	"encoding/json"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"myapp/models"
	"myapp/service"
	"path/filepath"
)

type Memes struct {
	s3Storage service.AwsS3Storage
	logger    *golog.Logger
}

func NewMemes(logger *golog.Logger, s3Storage service.AwsS3Storage) Memes {
	logger.Info("init Memes handle")
	return Memes{s3Storage: s3Storage, logger: logger}
}

func (m *Memes) GetMemes(ctx iris.Context) {
	memes := []models.Meme{
		{[]string{"11", "dddd"}, 1, []string{"slslsl", "olol"}, []string{"Tag"}, []string{"phrase"}},
		{[]string{"11", "dddd"}, 1, []string{"slslsl", "olol"}, []string{"Tag"}, []string{"phrase"}},
		{[]string{"11", "dddd"}, 1, []string{"slslsl", "olol"}, []string{"Tag"}, []string{"phrase"}},
	}
	ctx.JSON(memes)
}

func (m *Memes) PostMemes(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("failed to upload file").DetailErr(err))
		return
	}
	value := ctx.FormValue("meta")
	var mMeme models.Meme
	err = json.Unmarshal([]byte(value), &mMeme)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("failed to marshall meta data").DetailErr(err))
		return
	}
	m.logger.Debug("Meme marshalled succeed:", &mMeme)
	m.logger.Info("start uploading meme pic")
	picUrl, err := m.s3Storage.UploadMemePic(file, info.Filename, filepath.Ext(info.Filename))
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("meme upload failed,try it later?").DetailErr(err))
		return
	}
	m.logger.Info("meme pic upload finished url:", picUrl)

	ctx.StatusCode(iris.StatusCreated)
}
