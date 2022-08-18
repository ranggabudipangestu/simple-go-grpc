package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggabudipangestu/go-grpc-client/models"
	"github.com/ranggabudipangestu/go-grpc-client/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	userClient models.UsersClient
}

func CreateUserHandler(r *gin.Engine, userClient models.UsersClient) {
	userHandler := &UserHandler{userClient}

	r.POST("/user", userHandler.addUser)
	r.GET("/users", userHandler.viewUser)
	r.GET("/user/:id", userHandler.viewUserById)
	r.PUT("/user/:id", userHandler.updateUser)
	r.DELETE("/user/:id", userHandler.deleteUser)

}

func (handler *UserHandler) addUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = handler.userClient.InsertUser(context.Background(), &user)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, "Sucess Add Data")

}

func (handler *UserHandler) viewUser(c *gin.Context) {
	userList, err := handler.userClient.GetUserList(context.Background(), new(emptypb.Empty))
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, userList.List)
}

func (handler *UserHandler) viewUserById(c *gin.Context) {
	id := c.Param("id")
	userId := models.UserId{
		Id: id,
	}
	user, err := handler.userClient.GetUserById(context.Background(), &userId)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "failed to view data")
		return
	}

	utils.HandleSuccess(c, user)
}

func (handler *UserHandler) updateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := c.Bind(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to update User")
	}

	update := models.UserUpdate{
		Id:   id,
		User: &user,
	}

	_, err = handler.userClient.UpdateUser(context.Background(), &update)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to update User")
	}

	utils.HandleSuccess(c, "success update user")

}

func (handler *UserHandler) deleteUser(c *gin.Context) {
	id := c.Param("id")
	delete := models.UserId{
		Id: id,
	}
	_, err := handler.userClient.DeleteUser(context.Background(), &delete)

	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "data not found")
	}

	utils.HandleSuccess(c, "Success delete data")
}
