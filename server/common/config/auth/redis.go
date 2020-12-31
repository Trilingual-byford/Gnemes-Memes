package auth

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"time"
)

type RedisManagerOfAuth struct {
	Database *redis.Database
	Addr     string
	Logger   *golog.Logger
}

func Init(logger *golog.Logger) RedisManagerOfAuth {
	addr := "127.0.0.1:6379"
	redis := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      addr,
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Username:  "",
		Password:  "",
		Database:  "",
		Prefix:    "Gnemes-Auth",
		Driver:    redis.GoRedis(),
	})
	logger.Info("Init redis database successfully")

	return RedisManagerOfAuth{Logger: logger, Addr: addr, Database: redis}
}

func (manager RedisManagerOfAuth) SetAuthInfo() {

}
