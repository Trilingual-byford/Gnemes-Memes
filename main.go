package main

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"myapp/handles"
)

func main() {
	app := iris.New()
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "/doc/apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.
	app.Use(recover.New())
	app.Use(logger.New())
	memesApi := app.Party("/api/v1/gnemes")
	{
		memes := handles.Memes{}
		memesApi.Use(iris.Compression)
		memesApi.Get("/memes", memes.GetMemes)
		memesApi.Post("/memes", memes.PostMemes)
	}

	app.Listen(":8080")

}
