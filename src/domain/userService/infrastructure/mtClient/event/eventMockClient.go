package mtClient

import (
	"time"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	event "github.com/MindTickle/platform-protos/pb/event"
	"golang.org/x/net/context"
)

type EventMockClient struct {
}

func (client EventMockClient) CreateEvents(ctx context.Context, url string, channelId int64, eventData model.EventField) (*event.CreateEventsResponse, error) {
	event1 := event.Event{
		Id:         1,
		Authorizer: "DummyAuth",
		Source:     "DummySource",
		TenantId:   "DummyTenant",
		Data:       []byte(eventData.Data),
		Encoding:   event.Encoding_JSON,
		CreateTime: time.Now().Unix(),
	}

	res := event.CreateEventsResponse{}
	res.Events = append(res.Events, &event1)

	return &res, nil
}
