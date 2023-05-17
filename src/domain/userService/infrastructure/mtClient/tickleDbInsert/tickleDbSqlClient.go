package mtClient

import (
	"context"
	"fmt"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	pb "github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"google.golang.org/grpc"
)

type TickleDbSqlImpl struct {
	TickleDbSqlService pb.SqlStoreClient
	logger             *logger.LoggerImpl
}

func NewTickleDbSqlClient(l *logger.LoggerImpl) outgoing.TickleDbInsertClient {
	if constants.IsMock {
		return TickleDbSqlClient{}
	}

	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	res := TickleDbSqlImpl{
		TickleDbSqlService: pb.NewSqlStoreClient(conn),
		logger:             l,
	}

	return res
}

func (client TickleDbSqlImpl) InsertRow(id string, field model.User, url string, tableName string, reqContext pb.RequestContext, authMeta pb.AuthMeta) (*pb.CreateRowsResponse, error) {
	rowValue := pb.RowValue{
		RowInBytes: []byte(fmt.Sprintf(`{"uid": "%s", "uname":"%s"}`, field.UserId, field.UserName)),
		AuthMeta:   &authMeta,
	}

	row := pb.Row{
		Id:       id,
		RowValue: &rowValue,
	}
	data := pb.CreateRowsRequest{TableName: tableName, RequestContext: &reqContext}
	data.Rows = append(data.Rows, &row)
	res, err := client.TickleDbSqlService.CreateRows(context.Background(), &data)

	if err != nil {
		client.logger.Errorf(context.Background(), "TickleDbSqlStore Client Failed")
		return nil, err
	}

	return res, nil
}
