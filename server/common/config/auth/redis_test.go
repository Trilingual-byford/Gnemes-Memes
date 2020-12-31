package auth

import (
	"github.com/kataras/golog"
	"github.com/smartystreets/assertions"
	"testing"
	"time"
)

func TestRedisSetFunctionality(t *testing.T) {
	logger := golog.New()
	manager := Init(logger)
	db := manager.Database
	err := db.Set("niconicocsc", "TestRedisSetFunctionality", "2eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2VtYWlsIjoibmljb25pY29jc2NAZ21haWwuY29tIiwicm9sZXMiOlsidXNlciJdLCJpYXQiOjE2MDk0Mzg4MTQsImV4cCI6MTYwOTQzOTcxNCwianRpIjoiYmNiZmQwNTQtMTY5Yi00M2YwLWFkZDgtOTZhYmNkNmVmNmFiIiwiaXNzIjoiR25lbWVzIn0.cuOkMSJuhBimg4vLAzkV-u2kYJ0FdygKIq2Ax7qAYqg", time.Duration(30)*time.Hour, true)
	assertions.ShouldBeNil(err)
	value := db.Get("niconicocsc", "TestRedisSetFunctionality")
	//compare := strings.Compare("This is a key from TestRedisSetFunctionality func", fmt.Sprintf("%v",value))
	//compare3 := strings.Compare("This is a key from TestRedisSetFunctionality func", value.(string))
	//compare2 := strings.Compare("TestRedisSetFunctionality", "TestRedisSetFunctionality")
	assertions.ShouldNotBeEmpty(value)
	//assertions.ShouldNotBeEmpty(compare2)
	//assertions.ShouldNotBeEmpty(compare3)
}
func TestRedisGetFunctionality(t *testing.T) {
	logger := golog.New()
	manager := Init(logger)
	db := manager.Database
	value := db.Get("niconicocsc", "TestRedisSetFunctionality")
	assertions.ShouldNotBeEmpty(value)
}
