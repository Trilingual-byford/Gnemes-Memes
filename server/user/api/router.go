package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"gnemes/user/repository"
)

func NewRouter(repo repository.UserRepository, db *redis.Database) func(iris.Party) {
	return func(router iris.Party) {
		router.Post("/user/sign-in", SignIn(repo, db))
		//router.Post("/user/sign-out", SignOut())
		router.Post("/user/sign-up", SignUp(repo))
		router.Use(Verify())
		router.Get("/user/save/collection", func(ctx iris.Context) {
			ctx.JSON("collection")
		})
		router.Get("/user/like/collection")
		router.Get("/user/info")
		router.Get("/user/app/preference")

	}

}
