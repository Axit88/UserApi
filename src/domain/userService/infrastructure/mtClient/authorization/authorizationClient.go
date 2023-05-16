package main

import (
	"context"
	"fmt"

	"github.com/Axit88/UserApi/src/constants"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	pb "github.com/MindTickle/content-protos/pb/authorisation"
	"github.com/MindTickle/content-protos/pb/common"
	"github.com/MindTickle/mt-go-logger/logger"
	"google.golang.org/grpc"
)

type AuthorizationImpl struct {
	AuthorizationService pb.RolePermissionServiceClient
	logger               *logger.LoggerImpl
}

func NewAuthorizationClient(l *logger.LoggerImpl) outgoing.AuthorizationClient {

	if constants.IsMock {
		return AuthorizationMockClient{}
	}

	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	res := AuthorizationImpl{}
	res.AuthorizationService = pb.NewRolePermissionServiceClient(conn)
	res.logger = l
	return res
}

func (client AuthorizationImpl) GetCompnanyRolePermission(url string, companyId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) error {
	data := pb.GetRoleAndPermissionsRequest{RequestMeta: &reqMeta, RecordMeta: &recMeta}
	data.CompanyIds = append(data.CompanyIds, companyId)

	roles, err := client.AuthorizationService.GetRolesAndPermissions(context.Background(), &data)
	if err != nil {
		client.logger.Errorf(context.Background(), "Authorization Client Failed", err)
		return err
	}
	fmt.Println(roles)

	return nil
}

func (client AuthorizationImpl) GetUserRolePermission(url string, userId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) error {

	data := pb.GetUserRolesAndPermissionsRequest{UserId: userId, RequestMeta: &reqMeta, RecordMeta: &recMeta}

	roles, err := client.AuthorizationService.GetUserRolesAndPermissions(context.Background(), &data)
	if err != nil {
		client.logger.Errorf(context.Background(), "Authorization Client Failed", err)
		return err
	}
	fmt.Println(roles)

	return nil
}
