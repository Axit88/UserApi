package incoming

import "github.com/Axit88/UserApi/src/domain/userService/core/model"

type APIPort interface {
	ProcessAddUser(input *model.User) error
	ProcessGetUser(userId string) (*model.User, error)
	ProcessDeleteUser(userId string) error
	ProcessUpdateUser(userId string, userName string) error
}

type UserService interface {
	AddUser(input *model.User) error
	GetUser(userId string) (*model.User, error)
	DeleteUser(userId string) error
	UpdateUser(userId string, userName string) error
}
