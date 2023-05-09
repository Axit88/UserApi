package core

import (
	"testing"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/dbClient"
	"github.com/stretchr/testify/assert"
)

var facade *UserServiceImpl

func init() {
	constants.IsMock = true

	db, _ := dbClient.NewDbClient()
	facade = &UserServiceImpl{
		db: db,
	}

	constants.IsMock = false
}

func TestAddUser(t *testing.T) {

	userId := "1"
	userName := "Jay"
	input := adapters.GetCreateUserRequest(userId, userName)
	err := facade.AddUser(input)
	assert.Nil(t, err)

}

func TestGetUser(t *testing.T) {

	userId := "1"
	res, err := facade.GetUser(userId)
	assert.NotNil(t, res)
	assert.Nil(t, err)

}

func TestUpdateUser(t *testing.T) {

	userId := "1"
	userName := "Jay"
	err := facade.UpdateUser(userId, userName)
	assert.Nil(t, err)

}

func TestDeleteUser(t *testing.T) {

	userId := "1"
	err := facade.DeleteUser(userId)
	assert.Nil(t, err)

}
