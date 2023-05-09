package incoming

import "github.com/Axit88/UserApi/src/domain/userService/core/model"

type UserService interface {
	AddUser(input *model.User) error
	GetUser(userId string) (*model.User, error)
	DeleteUser(userId string) error
	UpdateUser(userId string, userName string) error
}
