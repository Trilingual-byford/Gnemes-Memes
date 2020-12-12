package main

import (
	"github.com/kataras/iris/v12"
	"gnemes/auth/api"
	"gnemes/auth/repository"
)

var (
	sigKey = []byte("signature_hmac_secret_shared_key")
)

func main() {
	app := iris.New()
	loggger := app.Logger()
	userRepo := repository.NewMongoUserRepository(loggger)
	app.PartyFunc("/", api.NewRouter(userRepo))
	app.Listen(":8082")
}
