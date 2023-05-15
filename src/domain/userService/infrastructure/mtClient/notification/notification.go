package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
)

type NotificationImpl struct {
	logger *logger.LoggerImpl
}

func NewNotificationClient(l *logger.LoggerImpl) outgoing.NotficationClient {

	if constants.IsMock {
		return NotificationMockClient{}
	}

	res := NotificationImpl{}
	res.logger = l
	return res
}

func (client NotificationImpl) SendNotification(url string, input model.NotificationField) (*model.EmailResponse, error) {
	res := model.EmailResponse{
		JobId: "",
	}
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	data := map[string]interface{}{
		"params": map[string]interface{}{
			"from_name":   input.From_Name,
			"domain_base": input.DomainBase,
			"cname":       input.Cname,
			"series":      input.Series,
			"entity":      input.Entity,
			"user_type":   input.UserType,
		},
		"reply_to":            input.ReplyTo,
		"from":                input.From,
		"mailer":              "ses",
		"category":            input.Category,
		"template_evaluation": "strict",
		"html":                true,
		"notificationChannel": input.NotificationChannel,
		"to":                  input.To,
		"template":            input.Template,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		client.logger.Errorf(context.Background(), "Payload Is Not Correct", err)
		return nil, err
	}

	url = fmt.Sprintf("%v/immediate", url)

	resp, err := c.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		client.logger.Errorf(context.Background(), "Send Notification Failed", err)
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		client.logger.Errorf(context.Background(), "Notification Response Is Not Correct", err)
		return nil, err
	}
	
	return &res, err
}
