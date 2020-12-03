package main

import (
	"github.com/kataras/iris/v12"
)

var (
	sigKey = []byte("signature_hmac_secret_shared_key")
)

func main() {
	app := iris.New()
	app.PartyFunc("/")

}
