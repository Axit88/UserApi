package dbClient

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"
)

type DbMockClient struct {
}

func (s DbMockClient) Insert(input *model.User) error {
	return nil
}

func (s DbMockClient) Update(userId string, userName string) error {
	return nil
}

func (s DbMockClient) Select(userId string) (*model.User, error) {
	return adapters.GetCreateUserRequest("99", "Tushar"), nil
}

func (s DbMockClient) Delete(userId string) error {
	return nil
}

func (s DbMockClient) CloseDbConnection() {
}
