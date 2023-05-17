package container

import (
	application "github.com/Axit88/UserApi/src/domain/userService/application"
	core "github.com/Axit88/UserApi/src/domain/userService/core"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/dbClient"
	gRPC "github.com/Axit88/UserApi/src/domain/userService/infrastructure/grpcClient"
	rest "github.com/Axit88/UserApi/src/domain/userService/infrastructure/restClient"
	"github.com/Axit88/UserApi/src/utils/loggerUtil"
	"go.uber.org/dig"
)

func UserServiceContainer() (*dig.Container, error) {

	container := dig.New()
	err := container.Provide(loggerUtil.InitLogger)
	err = container.Provide(dbClient.NewDbClient)
	err = container.Provide(application.NewApplication)
	err = container.Provide(core.NewFacadeClient)
	err = container.Provide(rest.NewHTTPHandler)
	err = container.Provide(gRPC.NewGrpcClient)

	return container, err
}
