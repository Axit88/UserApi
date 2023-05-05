package main

import (
	"log"

	"github.com/Axit88/UserApi/src/domain/userService/adapters/outgoing/db"
	application "github.com/Axit88/UserApi/src/domain/userService/application"
	core "github.com/Axit88/UserApi/src/domain/userService/core"
	"github.com/gin-gonic/gin"
	rest "github.com/Axit88/UserApi/src/domain/userService/adapters/incoming/rest"
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
	httpHandler := rest.NewHTTPHandler(applicationAPI)

	router := gin.Default()
	router.GET("/User/:id", httpHandler.GetUser)
	router.POST("/User", httpHandler.AddUser)
	router.PUT("/User/:id", httpHandler.UpdateUser)
	router.DELETE("/User/:id", httpHandler.DeleteUser)
	router.Run("localhost:9090")
}


// go run restServer.go