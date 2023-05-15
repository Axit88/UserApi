package main

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"golang.org/x/net/context"
)

type EventChannelMockClient struct {
}

func (client EventChannelMockClient) CreateEventChannel(ctx context.Context, url string, channelData model.EventChannelField) error {
	return nil
}
