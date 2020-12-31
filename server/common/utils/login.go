package utils

import (
	"errors"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"strings"
)

func CheckLoginStatus(redis *redis.Database, authSid string, userId string, token string) (isLogin bool, err error) {
	value := redis.Get(authSid, userId)
	if value == nil {
		return false, errors.New("authError:didn't login")
	}
	compareResult := strings.Compare(token, value.(string))
	switch compareResult {
	case 0:
		return true, nil
	default:
		return false, errors.New("token don't match")
	}
}
