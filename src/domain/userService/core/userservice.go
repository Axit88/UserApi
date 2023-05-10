package core

import (
	"context"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
)

type UserServiceImpl struct {
	logger *logger.LoggerImpl
	db     outgoing.DbPort
}

func New(db outgoing.DbPort, l *logger.LoggerImpl) incoming.UserService {
	return &UserServiceImpl{db: db, logger: l}
}

func (worker UserServiceImpl) AddUser(input *model.User) error {
	err := worker.db.Insert(input)
	if err != nil {
		worker.logger.Errorf(context.Background(), "Db Insert Failed", err)
	}

	return err
}

func (worker UserServiceImpl) GetUser(userId string) (*model.User, error) {
	res, err := worker.db.Select(userId)
	if err != nil {
		worker.logger.Errorf(context.Background(), "Db Select Failed", err)
	}

	return res, err
}

func (worker UserServiceImpl) DeleteUser(userId string) error {
	err := worker.db.Delete(userId)
	if err != nil {
		worker.logger.Errorf(context.Background(), "Db Delete Failed", err)
	}

	return err
}

func (worker UserServiceImpl) UpdateUser(userId string, userName string) error {
	err := worker.db.Update(userId, userName)
	if err != nil {
		worker.logger.Errorf(context.Background(), "Db Update Failed", err)
	}

	return err
}
