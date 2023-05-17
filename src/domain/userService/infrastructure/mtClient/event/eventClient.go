package mtClient

import (
	"fmt"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	event "github.com/MindTickle/platform-protos/pb/event"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EventImpl struct {
	EventService event.EventServiceClient
	logger       *logger.LoggerImpl
}

func NewEventClient(l *logger.LoggerImpl) outgoing.EventClient {

	if constants.IsMock {
		return EventMockClient{}
	}

	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	res := EventImpl{}
	res.EventService = event.NewEventServiceClient(conn)
	res.logger = l
	return res
}

func (client EventImpl) CreateEvents(ctx context.Context, url string, channelId int64, eventData model.EventField) (*event.CreateEventsResponse, error) {
	eventDetail := &event.Event{
		Data:       []byte(eventData.Data),
		Encoding:   event.Encoding_JSON,
		TenantId:   eventData.TenantId,
		Source:     eventData.Source,
		Authorizer: eventData.Authorizer,
	}
	data := event.CreateEventsRequest{EventChannelId: channelId}
	data.Events = append(data.Events, eventDetail)

	eventt, err := client.EventService.CreateEvents(ctx, &data)
	if err != nil {
		logger.Logger.Errorf(context.Background(), "Event Client Failed", err)
		return nil, err
	}

	return eventt, nil
}
