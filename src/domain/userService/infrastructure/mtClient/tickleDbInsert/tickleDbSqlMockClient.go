package mtClient

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	pb "github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
)

type TickleDbSqlClient struct {
}

func (client TickleDbSqlClient) InsertRow(id string, field model.User, url string, tableName string, reqContext pb.RequestContext, authMeta pb.AuthMeta) (*pb.CreateRowsResponse, error) {
	res := pb.CreateRowsResponse{
		RowsAffected: 1,
	}
	return &res, nil
}
