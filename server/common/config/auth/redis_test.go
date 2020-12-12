package auth

import (
	"github.com/kataras/golog"
	"github.com/smartystreets/assertions"
	"testing"
	"time"
)

func TestRedisSetFunctionality(t *testing.T) {
	logger := golog.New()
	db := GetRedisDatabase(logger)
	err := db.Set("niconicocsc", "TestRedisSetFunctionality", "This is a key from TestRedisSetFunctionality func", time.Duration(30)*time.Hour, true)
	assertions.ShouldBeNil(err)
	value := db.Get("niconicocsc", "TestRedisSetFunctionality")
	assertions.ShouldNotBeEmpty(value)
}
func TestRedisGetFuncionality(t *testing.T) {
	logger := golog.New()
	db := GetRedisDatabase(logger)
	value := db.Get("niconicocsc", "TestRedisSetFunctionality")
	assertions.ShouldNotBeEmpty(value)
}
