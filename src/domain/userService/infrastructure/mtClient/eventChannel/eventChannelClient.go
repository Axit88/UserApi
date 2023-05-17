package mtClient

import (
	"fmt"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	pb "github.com/MindTickle/platform-protos/pb/event_channel"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EventChannelImpl struct {
	EventChannelService pb.EventChannelServiceClient
	logger              *logger.LoggerImpl
}

func NewEventChannelClient(l *logger.LoggerImpl) outgoing.EventChannelClient {
	if constants.IsMock {
		return EventChannelMockClient{}
	}

	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	res := EventChannelImpl{
		EventChannelService: pb.NewEventChannelServiceClient(conn),
		logger:              l,
	}

	return res
}

func (client EventChannelImpl) CreateEventChannel(ctx context.Context, url string, channelData model.EventChannelField) error {

	newChannel := &pb.EventChannel{
		Name:        channelData.Name,
		Project:     channelData.Project,
		Parallelism: channelData.Parallelism,
		State:       pb.EventChannel_ENABLED,
	}
	_, err := client.EventChannelService.CreateEventChannel(ctx, &pb.CreateEventChannelRequest{EventChannel: newChannel})
	if err != nil {
		client.logger.Errorf(context.Background(), "EventChannel Client Failed", err)
		return err
	}

	return nil
}
