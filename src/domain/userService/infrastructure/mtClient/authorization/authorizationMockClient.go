package mtClient

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	pb "github.com/MindTickle/content-protos/pb/authorisation"
	"github.com/MindTickle/content-protos/pb/common"
)

type AuthorizationMockClient struct {
}

func NewAuthorizationMockClient() outgoing.AuthorizationClient {
	return &AuthorizationMockClient{}
}

func (client AuthorizationMockClient) GetCompnanyRolePermission(url string, companyId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) (*pb.GetRoleAndPermissionsResponse, error) {

	role := pb.RoleAndPermissionsDetails{
		Role: &pb.Role{
			RoleId:  "1",
			Company: "969694858385866",
			Name:    "DummyRole",
		},
	}

	company_role := pb.CompanyRoleAndPermissionsDetails{
		CompanyId: "999999999999",
	}
	company_role.RoleAndPermissionDetails = append(company_role.RoleAndPermissionDetails, &role)

	res := pb.GetRoleAndPermissionsResponse{}
	res.CompanyRoleAndPermissionsDetails = append(res.CompanyRoleAndPermissionsDetails, &company_role)

	return &res, nil
}

func (client AuthorizationMockClient) GetUserRolePermission(url string, userId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) (*pb.GetUserRolesAndPermissionsResponse, error) {
	role1 := pb.RoleAndPermissionsDetails{
		Role: &pb.Role{
			RoleId:  "1",
			Company: "969694858385866",
			Name:    "DummyRole",
		},
	}

	res := pb.GetUserRolesAndPermissionsResponse{
		UserId: "897648456754373",
	}
	res.RoleAndPermissionDetails = append(res.RoleAndPermissionDetails, &role1)

	return &res, nil
}
