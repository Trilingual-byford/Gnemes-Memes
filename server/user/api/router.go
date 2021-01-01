package api

import (
	"github.com/kataras/iris/v12"
	"gnemes/common/config/auth"
	"gnemes/user/api/handler"
	"gnemes/user/repository"
)

func NewRouter(repo repository.UserRepository, db auth.RedisManagerOfAuth) func(iris.Party) {
	return func(router iris.Party) {
		router.Post("/user/sign-in", handler.SignIn(repo, db.Database))
		//router.Post("/user/sign-out", SignOut())
		router.Post("/user/sign-up", handler.SignUp(repo))
		router.Use(handler.Verify(db.Database))
		idParty := router.Party("/id/{userId:string}")
		idParty.Get("/collections")
		idParty.Post("/collection")
		idParty.Delete("/collection/{gnemesId:string}")
		idParty.Get("/likes")
		idParty.Get("/info")
		idParty.Get("/app/preference")

		mailAddrParty := router.Party("/mail/{mailAddr:string}")
		mailAddrParty.Get("/collections")
		mailAddrParty.Get("/likes")
		mailAddrParty.Get("info")
		mailAddrParty.Get("/app/preference")
	}

}
