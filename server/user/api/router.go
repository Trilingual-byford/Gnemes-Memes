package api

import (
	"github.com/kataras/iris/v12"
	"gnemes/common/config/auth"
	"gnemes/user/repository"
)

func NewRouter(repo repository.UserRepository, db auth.RedisManagerOfAuth) func(iris.Party) {
	return func(router iris.Party) {
		router.Post("/user/sign-in", SignIn(repo, db.Database))
		//router.Post("/user/sign-out", SignOut())
		router.Post("/user/sign-up", SignUp(repo))
		router.Use(Verify(db.Database))
		router.Get("/{userId:string}/save/collection", func(ctx iris.Context) {
			ctx.JSON("collection")
		})
		router.Get("/{userId:string}/like/collection")
		router.Get("/{userId:string}/info")
		router.Get("/{userId:string}/app/preference")

	}

}
