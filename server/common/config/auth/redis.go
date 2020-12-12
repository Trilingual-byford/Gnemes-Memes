package auth

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"time"
)

func GetRedisDatabase(logger *golog.Logger) *redis.Database {
	redis := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      "127.0.0.1:6379",
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Username:  "",
		Password:  "",
		Database:  "",
		Prefix:    "Gnemes-Auth",
		Driver:    redis.GoRedis(),
	})
	logger.Info("Init redis database successfully")

	logger.Info("Init redis session successfully")
	return redis
}
