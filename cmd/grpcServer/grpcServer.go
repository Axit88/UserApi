package main

import (
	"log"

	gRPC "github.com/Axit88/UserApi/src/domain/userService/adapters/incoming/grpc"
	"github.com/Axit88/UserApi/src/domain/userService/adapters/outgoing/db"
	application "github.com/Axit88/UserApi/src/domain/userService/application"
	core "github.com/Axit88/UserApi/src/domain/userService/core"
)

func main() {
	var err error

	dbaseDriver := "mysql"
	dsourceName := "root:root@tcp(localhost:3306)/UserService"

	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core := core.New(dbAdapter)
	applicationAPI := application.NewApplication(core)
	gRPCAdapter := gRPC.NewAdapter(applicationAPI)
	gRPCAdapter.Run()
}

// go run grpcServer.go
