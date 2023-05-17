package mtClient

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	pb "github.com/MindTickle/storageprotos/pb/tickleDb"
)

type TickleDbStoreMockClient struct {
}

func (client TickleDbStoreMockClient) CreateTable(dbDetail model.TickleDbEnvDetail) (*pb.CreateTableResponse, error) {
	res := pb.CreateTableResponse{
		TableName: "dummy",
	}
	return &res, nil
}
