package mtClient

import (
	"time"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"
)

type AuthenticationMockClient struct {
}

func (client AuthenticationMockClient) VerifySession(url string, sessionId string) (*model.AuthenticatioResponse, error) {
	res := model.AuthenticatioResponse{
		Id:         "123",
		Email:      "dummy@gmail.com",
		Name:       "Jessy",
		OrgId:      "999999990999",
		Timezone:   time.Hour.Nanoseconds(),
		CompanyId:  "888888888888",
		SessionKey: "SSKg78FDFFFSsd5swwd",
	}

	return &res, nil
}
