package handles

import (
	"github.com/kataras/iris/v12"
	"myapp/models"
)

type Memes struct {}


func (m *Memes)GetMemes(ctx iris.Context){
	memes := []models.Meme{
		{[]string{"11","dddd"},1,[]string{"slslsl","olol"},[]string{"Tag"},[]string{"phrase"}},
		{[]string{"11","dddd"},1,[]string{"slslsl","olol"},[]string{"Tag"},[]string{"phrase"}},
		{[]string{"11","dddd"},1,[]string{"slslsl","olol"},[]string{"Tag"},[]string{"phrase"}},
	}
	ctx.JSON(memes)
}

func (m *Memes)PostMemes(ctx iris.Context){
	var meme models.Meme
	err:= ctx.ReadJSON(&meme)
	if err!=nil{
		ctx.StopWithProblem(iris.StatusBadRequest,iris.NewProblem().Title("Meme creation failure").DetailErr(err))
		return
	}
	println("Received Meme:",meme.OLSentences)
	ctx.StatusCode(iris.StatusCreated)

}