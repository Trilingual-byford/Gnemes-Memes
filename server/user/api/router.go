package api

import (
	"github.com/kataras/iris/v12"
	"gnemes/user/repository"
)

func NewRouter(repo repository.UserRepository) func(iris.Party) {
	return func(router iris.Party) {
		router.Post("/signin", SignIn(repo))
		router.Post("/signup", SignUp(repo))
		router.Use(Verify())
	}

}
