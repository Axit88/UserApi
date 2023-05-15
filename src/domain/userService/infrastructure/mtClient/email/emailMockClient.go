package main

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
)

type EmailMockClient struct {
}

func (client EmailMockClient) SendEmail(url string, input model.EmailField) (*model.EmailResponse, error) {
	return nil, nil
}
