package main

import (
	"github.com/kataras/iris/v12"
	"gnemes/common/config/auth"
	"gnemes/user/api"
	"gnemes/user/repository"
)

var (
	sigKey = []byte("signature_hmac_secret_shared_key")
)

func main() {
	app := iris.New()
	logger := app.Logger()
	userRepo := repository.NewMongoUserRepository(logger)
	authDB := auth.GetRedisDatabase(logger)
	app.PartyFunc("/", api.NewRouter(userRepo, authDB))
	err := app.Listen(":8082")
	if err != nil {
		//TODO
	}
}
