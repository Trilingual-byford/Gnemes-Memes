package auth

import (
	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGETALLUSER(t *testing.T) {
	logger := golog.New()
	sessions := GetRedisSessions(logger)
	sessions.
		assert.NotEmpty(t, all)
}
