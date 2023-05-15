package api

import (
	"context"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	incoming "github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"
	"github.com/MindTickle/mt-go-logger/logger"
)

// Application implements the APIPort interface
type Application struct {
	logger *logger.LoggerImpl
	facade incoming.UserService
}

func NewApplication(usr incoming.UserService, l *logger.LoggerImpl) incoming.APIPort {
	return &Application{facade: usr, logger: l}
}

func (apia Application) ProcessAddUser(userId string, userName string) error {
	input := adapters.GetCreateUserRequest(userId, userName)

	err := apia.facade.AddUser(input)
	if err != nil {
		apia.logger.Errorf(context.Background(), "Failed To Process Add User Application Request", err)
	}

	return err
}

func (apia Application) ProcessGetUser(userId string) (*model.User, error) {

	res, err := apia.facade.GetUser(userId)
	if err != nil {
		apia.logger.Errorf(context.Background(), "Failed To Process Get User Application Request", err)
	}

	return res, err
}

func (apia Application) ProcessDeleteUser(userId string) error {
	err := apia.facade.DeleteUser(userId)
	if err != nil {
		apia.logger.Errorf(context.Background(), "Failed To Process Delete User Application Request", err)
	}

	return err
}

func (apia Application) ProcessUpdateUser(userId string, userName string) error {
	err := apia.facade.UpdateUser(userId, userName)
	if err != nil {
		apia.logger.Errorf(context.Background(), "Failed To Process Update User Application Request", err)
	}

	return err
}
