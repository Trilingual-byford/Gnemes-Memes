package auth

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"time"
)

func GetRedisSessions(logger *golog.Logger) *sessions.Sessions {
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
	// Optionally configure the underline driver
	// driver := redis.Redis()
	// driver.ClientOptions = redis.Options{...}
	// driver.ClusterOptions = redis.ClusterOptions{}
	// redis.New(redis.Config{Driver:driver,....})
	defer redis.Close()
	redis.Set()
	session := sessions.New(sessions.Config{
		Cookie:          "_session_id",
		Expires:         0,
		AllowReclaim:    true, //default to 0: unlimited life.Another good value is:45 * time.Minute,
		CookieSecureTLS: true,
	})
	session.UseDatabase(redis)
	logger.Info("Init redis session successfully")
	return session
}
