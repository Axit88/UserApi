package main

import (
	"github.com/Axit88/UserApi/src/container"
	gRPC "github.com/Axit88/UserApi/src/domain/userService/infrastructure/grpcClient"
	"github.com/Axit88/UserApi/src/utils"
)

func main() {

	utils.SetEnv()

	cont, err := container.UserServiceContainer()
	if err != nil {
		panic(err)
	}

	err = cont.Invoke(gRPC.StartServer)
	if err != nil {
		panic(err)
	}
}

// go run grpcServer.go
// go run src/server/grpcServer/grpcServer.go
