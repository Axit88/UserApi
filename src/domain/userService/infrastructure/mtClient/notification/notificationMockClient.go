package main

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
)

type NotificationMockClient struct {
}

func (client NotificationMockClient) SendNotification(url string, input model.NotificationField) (*model.EmailResponse, error) {

	return nil, nil
}
