package adapters

import "github.com/Axit88/UserApi/src/domain/userService/core/model"

func GetCreateUserRequest(userId string, userName string) *model.User {
	return &model.User{
		UserId:   userId,
		UserName: userName,
	}
}
