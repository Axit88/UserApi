package core

import (
	"testing"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/dbClient"
	"github.com/Axit88/UserApi/src/utils/loggerUtil"
	"github.com/stretchr/testify/assert"
)

var facade *UserServiceImpl
var Cases []model.User

func init() {
	constants.IsMock = true
	l, _ := loggerUtil.InitLogger()
	db, _ := dbClient.NewDbClient(l)
	facade = &UserServiceImpl{
		db:     db,
		logger: l,
	}

	Cases = append(Cases, model.User{UserId: "1", UserName: "Abhi"})
	Cases = append(Cases, model.User{UserId: "2", UserName: "Raj"})
	Cases = append(Cases, model.User{UserId: "3", UserName: "Sumit"})
	Cases = append(Cases, model.User{UserId: "4", UserName: "Jay"})

	constants.IsMock = false
}

func TestAddUser(t *testing.T) {

	for _, usr := range Cases {
		input := adapters.GetCreateUserRequest(usr.UserId, usr.UserName)
		err := facade.AddUser(input)
		assert.Nil(t, err)
	}
}

func TestGetUser(t *testing.T) {

	for _, usr := range Cases {
		res, err := facade.GetUser(usr.UserId)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	}
}

func TestUpdateUser(t *testing.T) {

	for _, usr := range Cases {
		err := facade.UpdateUser(usr.UserId, usr.UserName)
		assert.Nil(t, err)
	}
}

func TestDeleteUser(t *testing.T) {

	for _, usr := range Cases {
		err := facade.DeleteUser(usr.UserId)
		assert.Nil(t, err)
	}
}
