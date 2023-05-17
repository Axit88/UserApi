package mtClient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
)

type AuthenticationImpl struct {
	logger *logger.LoggerImpl
}

func NewAuthenticationClient(l *logger.LoggerImpl) outgoing.AuthenticationClient {

	if constants.IsMock {
		return AuthenticationMockClient{}
	}

	res := AuthenticationImpl{}
	res.logger = l
	return res
}

func (client AuthenticationImpl) VerifySession(url string, sessionId string) (*model.AuthenticatioResponse, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return &model.AuthenticatioResponse{}, nil
	}

	req.Header.Set("x-token", sessionId)

	newClient := &http.Client{}
	resp, _ := newClient.Do(req)
	var data model.AuthenticatioResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return &model.AuthenticatioResponse{}, nil
	}

	return &data, nil
}
