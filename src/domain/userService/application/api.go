package api

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	incoming "github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"
)

// Application implements the APIPort interface
type Application struct {
	facade incoming.UserService
}

func NewApplication(usr incoming.UserService) incoming.APIPort {
	return &Application{facade: usr}
}

func (apia Application) ProcessAddUser(userId string, userName string) error {
	input := adapters.GetCreateUserRequest(userId, userName)
	return apia.facade.AddUser(input)
}

func (apia Application) ProcessGetUser(userId string) (*model.User, error) {
	return apia.facade.GetUser(userId)
}

func (apia Application) ProcessDeleteUser(userId string) error {
	return apia.facade.DeleteUser(userId)
}

func (apia Application) ProcessUpdateUser(userId string, userName string) error {
	return apia.facade.UpdateUser(userId, userName)
}
