package utils

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

func CheckLoginStatus(ctx *context.Context, redis *redis.Database, authSid string, verifiedTokenContextKey string) error {
	//verifiedToken := ctx.Values().Get(verifiedTokenContextKey)
	//verifiedToken.
	return nil
}
