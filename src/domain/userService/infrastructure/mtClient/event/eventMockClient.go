package main

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	event "github.com/MindTickle/platform-protos/pb/event"
	"golang.org/x/net/context"
)

type EventMockClient struct {
}

func (client EventMockClient) CreateEvents(ctx context.Context, url string, channelId int64, eventData model.EventField) (*event.CreateEventsResponse, error) {
	return nil, nil
}
