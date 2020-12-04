package main

import (
	"github.com/kataras/iris/v12"
	"gnemes/user/api"
	"gnemes/user/repository"
)

var (
	sigKey   = []byte("signature_hmac_secret_shared_key")
	userRepo = repository.NewMongoUserRepository()
)

func main() {
	app := iris.New()
	app.PartyFunc("/", api.NewRouter(userRepo))
	app.Listen(":8081")

}
