package api

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	incoming "github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	"github.com/gin-gonic/gin"
)

// Application implements the APIPort interface
type Application struct {
	usr incoming.UserService
	r   *gin.Engine
}

func NewApplication(usr incoming.UserService) *Application {
	return &Application{usr: usr}
}

func (apia Application) ProcessAddUser(input *model.User) error {
	return apia.usr.AddUser(input)
}

func (apia Application) ProcessGetUser(userId string) (*model.User, error) {
	return apia.usr.GetUser(userId)
}

func (apia Application) ProcessDeleteUser(userId string) error {
	return apia.usr.DeleteUser(userId)
}

func (apia Application) ProcessUpdateUser(userId string, userName string) error {
	return apia.usr.UpdateUser(userId, userName)
}
