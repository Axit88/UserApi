package main

import (
	"github.com/Axit88/UserApi/src/container"
	rest "github.com/Axit88/UserApi/src/domain/userService/infrastructure/restClient"
	"github.com/Axit88/UserApi/src/utils"
)

func main() {
	utils.SetEnv()

	cont, err := container.UserServiceContainer()
	if err != nil {
		panic(err)
	}

	err = cont.Invoke(rest.StartServer)
	if err != nil {
		panic(err)
	}
}


// go run src/server/restServer/restServer.go
