package outgoing

import "github.com/Axit88/UserApi/src/domain/userService/core/model"

type DbClient interface {
	CloseDbConnection()
	Insert(input *model.User) error
	Select(userId string) (*model.User, error)
	Delete(userId string) error
	Update(userId string, userName string) error
}
