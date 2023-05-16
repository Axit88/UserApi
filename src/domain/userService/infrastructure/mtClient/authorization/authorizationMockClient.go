package main

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/content-protos/pb/common"
)

type AuthorizationMockClient struct {
}

func NewAuthorizationMockClient() (outgoing.AuthorizationClient) {
	return &AuthorizationMockClient{}
}


func (client AuthorizationMockClient) GetCompnanyRolePermission(url string, companyId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) error {
	return nil
}

func (client AuthorizationMockClient) GetUserRolePermission(url string, userId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) error {
	return nil
}
