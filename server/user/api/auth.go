package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"gnemes/common/model"
	"gnemes/common/utils"
	"gnemes/user/repository"
	"os"
	"time"
)

const defaultSecretKey = "sercrethatmaycontainch@r$32chars"
const authSid = "Auth information of Gnemes User"

func getSecretKey() string {
	secret := os.Getenv(utils.AppName + "_SECRET")
	if secret == "" {
		return defaultSecretKey
	}

	return secret
}

// UserClaims represents the user token claims.
type UserClaims struct {
	UserEmail string       `json:"user_email"`
	Roles     []model.Role `json:"roles"`
}

// Validate implements the custom struct claims validator,
// this is totally optionally and maybe unnecessary but good to know how.
func (u *UserClaims) Validate() error {
	if u.UserEmail == "" {
		return fmt.Errorf("%w: %s", jwt.ErrMissingKey, "user_id")
	}

	return nil
}

// Verify allows only authorized clients.
func Verify() iris.Handler {
	secret := getSecretKey()
	verifier := jwt.NewVerifier(jwt.HS256, []byte(secret), jwt.Expected{Issuer: utils.AppName})
	verifier.Extractors = []jwt.TokenExtractor{jwt.FromHeader} // extract token only from Authorization: Bearer $token
	//return verifier.Verify(func() interface{} {
	//
	//	return new(UserClaims)
	//})
	return func(ctx *context.Context) {
		token := []byte(verifier.RequestToken(ctx))
		verifiedToken, err := verifier.VerifyToken(token)
		//verifiedToken.Payload
		if err != nil {

		}
		const verifiedTokenContextKey = "iris.jwt.token"
		ctx.Values().Set(verifiedTokenContextKey, verifiedToken)
		ctx.Next()
	}
}

// AllowAdmin allows only authorized clients with "admin" access role.
// Should be registered after Verify.
func AllowAdmin(ctx iris.Context) {
	if !IsAdmin(ctx) {
		ctx.StopWithText(iris.StatusForbidden, "admin access required")
		return
	}
	ctx.Next()
}

func SignUp(repo repository.UserRepository) iris.Handler {
	return func(ctx iris.Context) {
		var (
			pwd      = ctx.FormValue("password")
			username = ctx.FormValue("username")
			email    = ctx.FormValue("email")
			sex      = ctx.FormValue("sex")
		)
		hashedPassword, err := utils.GeneratePassword(pwd)
		if err != nil {
			ctx.StopWithJSON(iris.StatusBadRequest, err)
		}

		sexType, err := model.GetSexTypeFromString(sex)
		if err != nil {
			ctx.StopWithJSON(iris.StatusBadRequest, err)
		}
		user, err := repo.Create(username, string(hashedPassword), email, "", sexType)
		if err != nil {
			ctx.StopWithJSON(iris.StatusBadRequest, err)
		} else {
			ctx.JSON(user)
		}
	}
}

// SignIn accepts the user form data and returns a token to authorize a client.
func SignIn(repo repository.UserRepository, db *redis.Database) iris.Handler {
	secret := getSecretKey()
	signer := jwt.NewSigner(jwt.HS256, []byte(secret), 15*time.Minute)

	return func(ctx iris.Context) {
		/*
			type LoginForm struct {
				Username string `form:"username"`
				Password string `form:"password"`
			}
			and ctx.ReadForm OR use the ctx.FormValue(s) method.
		*/

		var (
			userEmail = ctx.FormValue("userEmail")
			password  = ctx.FormValue("password")
		)
		user, ok := repo.GetByUserEmailAndPassword(userEmail, password)
		if !ok {
			ctx.StopWithText(iris.StatusBadRequest, "wrong username or password")
			return
		}
		// Optionally, generate a JWT ID.
		jti, err := utils.GenerateUUID()
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}
		claims := UserClaims{
			UserEmail: user.Email,
			Roles:     user.Roles,
		}
		now := time.Now()
		expiresAt := now.Add(time.Minute * 60)

		token, err := signer.Sign(claims, jwt.Claims{
			ID:       jti,
			Issuer:   utils.AppName,
			IssuedAt: now.Unix(),
			Expiry:   expiresAt.Unix(),
		})
		redisErr := db.Set(authSid, user.Email, "token", time.Duration(30)*time.Second, true)
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, redisErr)
		}

		ctx.Write(token)
	}
}

// SignOut invalidates a user from server-side using the jwt Blocklist.
func SignOut(ctx iris.Context) iris.Handler {
	return func(ctx iris.Context) {
		ctx.Logout() // this is automatically binded to a function which invalidates the current request token by the JWT Verifier above.
	}
}

// GetClaims returns the current authorized client claims.
func GetClaims(ctx iris.Context) *UserClaims {
	claims := jwt.Get(ctx).(*UserClaims)
	return claims
}

// GetUserID returns the current authorized client's user id extracted from claims.
func GetUserEmail(ctx iris.Context) string {
	return GetClaims(ctx).UserEmail
}

// IsAdmin reports whether the current client has admin access.
func IsAdmin(ctx iris.Context) bool {
	for _, role := range GetClaims(ctx).Roles {
		if role == model.Admin {
			return true
		}
	}

	return false
}
