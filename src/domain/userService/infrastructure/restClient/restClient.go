package infrastructure

import (
	"github.com/Axit88/UserApi/src/config"
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"

	"net/http"

	//"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc incoming.APIPort
}

func NewHTTPHandler(userService incoming.APIPort) *HTTPHandler {
	return &HTTPHandler{
		svc: userService,
	}
}

func (h HTTPHandler) GetUser(context *gin.Context) {
	id := context.Param("id")
	res, err := h.svc.ProcessGetUser(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Core Is Not Able To Process Request": err})
		return
	}

	context.JSON(http.StatusOK, res)
}

func (h HTTPHandler) AddUser(context *gin.Context) {

	newUser := adapters.GetCreateUserRequest("", "")
	err := context.BindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json Payload"})
		return
	}

	err = h.svc.ProcessAddUser(newUser.UserId, newUser.UserName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Core Is Not Able To Process Request": err})
		return
	}

	context.IndentedJSON(http.StatusCreated, "User Added") // (status , JSON)
}

func (h HTTPHandler) DeleteUser(context *gin.Context) {
	id := context.Param("id")

	err := h.svc.ProcessDeleteUser(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Core Is Not Able To Process Reuqest": err})
		return
	}

	context.IndentedJSON(http.StatusOK, "Deleted Successfully")
}

func (h HTTPHandler) UpdateUser(context *gin.Context) {

	newUser := adapters.GetCreateUserRequest("", "")
	err := context.BindJSON(&newUser)
	if err != nil {
		return
	}

	id := context.Param("id")
	err = h.svc.ProcessUpdateUser(id, newUser.UserName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Core Is Not Able To Process Reuqest": err})
		return
	}

	context.IndentedJSON(http.StatusOK, "Updated Successfully")
}

func StartServer(handler *HTTPHandler) {
	router := gin.Default()
	router.GET("/User/:id", handler.GetUser)
	router.POST("/User", handler.AddUser)
	router.PUT("/User/:id", handler.UpdateUser)
	router.DELETE("/User/:id", handler.DeleteUser)

	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.RestUrl
	router.Run(url)
}