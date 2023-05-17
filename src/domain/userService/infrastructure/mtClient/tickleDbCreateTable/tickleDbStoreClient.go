package mtClient

import (
	"context"
	"fmt"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	pb "github.com/MindTickle/storageprotos/pb/tickleDb"
	"google.golang.org/grpc"
)

type TickleDbStoreImpl struct {
	TickleDbStoreImplService pb.StoreManagerClient
	logger                   *logger.LoggerImpl
}

func NewTickleDbStoreImplClient(l *logger.LoggerImpl) outgoing.TickleDbCreateTableClient {

	if constants.IsMock {
		return TickleDbStoreMockClient{}
	}

	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	res := TickleDbStoreImpl{}
	res.TickleDbStoreImplService = pb.NewStoreManagerClient(conn)
	res.logger = l

	return res
}

func (client TickleDbStoreImpl) CreateTable(dbDetail model.TickleDbEnvDetail) (*pb.CreateTableResponse, error) {

	myTable := &pb.Table{
		TableName: dbDetail.TableName,
		Ttl:       0,
		Version:   0,
		Namespace: dbDetail.Namespace,
		Env:       dbDetail.Env,
		Columns: []*pb.Field{
			&pb.Field{
				FieldName:          "uid",
				DataType:           1,
				EnumExpectedValues: nil,
				NestedFields:       nil,
				Required:           true,
				Size:               10,
				DefaultValue:       "1",
			},
			&pb.Field{
				FieldName:          "uname",
				DataType:           1,
				EnumExpectedValues: nil,
				NestedFields:       nil,
				Required:           true,
				Size:               255,
				DefaultValue:       "",
			},
		},
		PrimaryKey: &pb.PrimaryKey{Columns: []string{"uid"}},
		//IndexColumns:       []*pb.IndexField{{FieldPath: []string{"uname"}},},
		PartitionStrategy: pb.PartitionStrategy_HASH_BASED,
		PartitionKey:      "",
	}

	res, err := client.TickleDbStoreImplService.CreateTable(context.Background(), &pb.CreateTableRequest{Table: myTable})
	if err != nil {
		client.logger.Errorf(context.Background(), "TickleDbStore Client Failed", err)
		return nil, err
	}

	return res, nil
}
