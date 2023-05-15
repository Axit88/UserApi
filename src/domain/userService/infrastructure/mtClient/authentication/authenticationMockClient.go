package main

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
)

type AuthenticationMockClient struct {
}

func (client AuthenticationMockClient) VerifySession(url string, sessionId string) (*model.AuthenticatioResponse, error) {
	return nil, nil
}
