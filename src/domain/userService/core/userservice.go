package core

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
)

type UserServiceImpl struct {
	db outgoing.DbPort
}

func New(db outgoing.DbPort) incoming.UserService {
	return &UserServiceImpl{db: db}
}

func (worker UserServiceImpl) AddUser(input *model.User) error {
	return worker.db.Insert(input)
}

func (worker UserServiceImpl) GetUser(userId string) (*model.User, error) {
	return worker.db.Select(userId)
}

func (worker UserServiceImpl) DeleteUser(userId string) error {
	return worker.db.Delete(userId)
}

func (worker UserServiceImpl) UpdateUser(userId string, userName string) error {
	return worker.db.Update(userId, userName)
}
