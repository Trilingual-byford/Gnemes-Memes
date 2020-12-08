package repository

import (
	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGETALLUSER(t *testing.T) {
	logger := golog.New()
	userRepo := NewMongoUserRepository(logger)
	all, err := userRepo.GetAll()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.NotEmpty(t, all)
}
