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
func TestFindUERBYEMAILANDPASSWORD_POSITIVE(t *testing.T) {
	logger := golog.New()
	userRepo := NewMongoUserRepository(logger)
	user, b := userRepo.GetByUserEmailAndPassword("niconicocsc@gmail.com", "password")
	assert.True(t, b)
	assert.NotEmpty(t, user)
}
func TestFindUERBYEMAILANDPASSWORD_NEGATIVE(t *testing.T) {
	logger := golog.New()
	userRepo := NewMongoUserRepository(logger)
	user, b := userRepo.GetByUserEmailAndPassword("niconicocsc@gmail.com", "wrongpassword")
	assert.False(t, b)
	assert.NotEmpty(t, user)
}
