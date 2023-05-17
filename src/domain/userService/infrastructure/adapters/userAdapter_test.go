package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserRequest(t *testing.T) {
	userId := "1"
	userName := "Nil"
	res := GetUserRequest(userId, userName)

	assert.Equal(t, res.UserId, userId)
	assert.Equal(t, res.UserName, userName)
}
