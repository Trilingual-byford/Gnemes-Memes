package handles

import (
	"context"
	"encoding/json"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/mongo"
	"myapp/config"
	"myapp/models/db"
	"myapp/models/dto"
	"myapp/service"
	"path/filepath"
	"time"
)

type Memes struct {
	s3Storage      service.AwsS3Storage
	logger         *golog.Logger
	memeCollection *mongo.Collection
}

func NewMemes(logger *golog.Logger, s3Storage service.AwsS3Storage) Memes {
	logger.Info("init Memes handle")
	db, err := config.DB(logger)
	if err != nil {
		logger.Error("failed to initial mongo db")
	}
	collection := db.Database("gnemes").Collection("GnemesPost")

	return Memes{s3Storage: s3Storage, logger: logger, memeCollection: collection}
}

func (m *Memes) GetMemes(ctx iris.Context) {
	memes := []dto.Meme{
		{[]string{"11", "dddd"}, 1, []string{"slslsl", "olol"}, []string{"Tag"}, []string{"phrase"}},
		{[]string{"11", "dddd"}, 1, []string{"slslsl", "olol"}, []string{"Tag"}, []string{"phrase"}},
		{[]string{"11", "dddd"}, 1, []string{"slslsl", "olol"}, []string{"Tag"}, []string{"phrase"}},
	}
	ctx.JSON(memes)
}

func (m *Memes) PostMemes(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	defer file.Close()
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("failed to upload file").DetailErr(err))
		return
	}
	value := ctx.FormValue("meta")
	var mMeme dto.Meme
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
	gneme := db.Gneme{
		CreatedTime: time.Now(),
		Difficulty:  mMeme.Difficulty,
		Dir:         picUrl,
		Likes:       0,
		Viewer:      0,
		Tag:         mMeme.Tag,
		PicHash:     "TestTestTest",
		OLSentences: mMeme.OLSentences,
		SLSentences: mMeme.SLSentences,
		Phrase:      mMeme.Phrase,
	}

	result, err := m.memeCollection.InsertOne(
		context.Background(), gneme)
	if err != nil {
		m.logger.Error("failed to upload meme to db", err)
	}
	m.logger.Info("meme have been updated to db:%+v", result)

	m.logger.Info("meme pic upload finished url:", picUrl)

	ctx.StatusCode(iris.StatusCreated)
}
