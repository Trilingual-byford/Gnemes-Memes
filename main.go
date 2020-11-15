package main

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"myapp/handles"
	"myapp/service"
	"os"
)

func makeAccessLog() *accesslog.AccessLog {
	// Initialize a new access log middleware.
	ac := accesslog.File("./access.log")
	// Remove this line to disable logging to console:
	ac.AddOutput(os.Stdout)
	// The default configuration:
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	// Default line format if formatter is missing:
	// Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
	//
	// Set Custom Formatter:
	ac.SetFormatter(&accesslog.JSON{
		Indent:    "  ",
		HumanTime: true,
	})
	// ac.SetFormatter(&accesslog.CSV{})
	// ac.SetFormatter(&accesslog.Template{Text: "{{.Code}}"})

	return ac
}

func main() {
	accessLog := makeAccessLog()
	defer accessLog.Close()
	app := iris.New()
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "/doc/apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.
	app.Use(recover.New())
	// Register the middleware (UseRouter to catch http errors too).
	app.UseRouter(accessLog.Handler)
	loggger := app.Logger()
	memesApi := app.Party("/api/v1/gnemes")
	{
		storage := service.NewAwsS3Storage()
		memes := handles.NewMemes(loggger, storage)
		memesApi.Use(iris.Compression)
		memesApi.Get("/memes", memes.GetMemes).Describe("Get all memes")
		memesApi.Post("/memes", memes.PostMemes).Describe("Insert a Meme to system ")
	}

	app.Listen(":8080", iris.WithLogLevel("info"))
}
