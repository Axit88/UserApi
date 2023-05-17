package application

import (
	"testing"

	"github.com/Axit88/UserApi/src/constants"
	core "github.com/Axit88/UserApi/src/domain/userService/core"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/dbClient"
	"github.com/Axit88/UserApi/src/utils/loggerUtil"
	"github.com/stretchr/testify/assert"
)

var app *Application
var Cases []model.User

func init() {
	constants.IsMock = true
	l, _ := loggerUtil.InitLogger()
	db, _ := dbClient.NewDbClient(l)
	facade := core.NewFacadeClient(db, l)

	app = &Application{
		facade: facade,
		logger: l,
	}

	Cases = append(Cases, model.User{UserId: "1", UserName: "Abhi"})
	Cases = append(Cases, model.User{UserId: "2", UserName: "Raj"})
	Cases = append(Cases, model.User{UserId: "3", UserName: "Sumit"})
	Cases = append(Cases, model.User{UserId: "4", UserName: "Jay"})

	constants.IsMock = false
}

func TestProcessAddUser(t *testing.T) {

	for _, usr := range Cases {
		err := app.ProcessAddUser(usr.UserId, usr.UserName)
		assert.Nil(t, err)
	}
}

func TestProcessGetUser(t *testing.T) {

	for _, usr := range Cases {
		res, err := app.ProcessGetUser(usr.UserId)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	}
}

func TestProcessUpdateUser(t *testing.T) {

	for _, usr := range Cases {
		err := app.ProcessUpdateUser(usr.UserId, usr.UserName)
		assert.Nil(t, err)
	}
}

func TestProcessDeleteUser(t *testing.T) {

	for _, usr := range Cases {
		err := app.ProcessDeleteUser(usr.UserId)
		assert.Nil(t, err)
	}
}
