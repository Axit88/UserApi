package incoming

import "github.com/Axit88/UserApi/src/domain/userService/core/model"

type APIPort interface {
	ProcessAddUser(userId string, userName string) error
	ProcessGetUser(userId string) (*model.User, error)
	ProcessDeleteUser(userId string) error
	ProcessUpdateUser(userId string, userName string) error
}
