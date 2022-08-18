package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ranggabudipangestu/go-grpc-server/models"
	service "github.com/ranggabudipangestu/go-grpc-server/modules/user/service"
	"google.golang.org/grpc"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(grpc *grpc.Server, service service.UserService) {
	userHandler := UserHandler{service}
	models.RegisterUsersServer(grpc, &userHandler)
}

func (handler *UserHandler) GetUserList(ctx context.Context, in *empty.Empty) (*models.UserList, error) {
	users, err := handler.userService.FindUsers()
	if err != nil {
		return nil, err
	}
	var userChannel = make([]*models.User, 0)
	for i := 0; i < len(*users); i++ {
		data := new(models.User)
		data.Id = (*users)[i].Id
		data.Email = (*users)[i].Email
		data.Name = (*users)[i].Name
		data.Alamat = (*users)[i].Alamat
		data.Password = (*users)[i].Password
		userChannel = append(userChannel, data)
	}
	var u = models.UserList{
		List: userChannel,
	}

	return &u, nil
}

func (handler *UserHandler) GetUserById(ctx context.Context, userId *models.UserId) (*models.User, error) {
	user, err := handler.userService.FindUserById(userId)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (handler *UserHandler) InsertUser(ctx context.Context, user *models.User) (*empty.Empty, error) {
	_, err := handler.userService.AddUser(user)

	if err != nil {
		return nil, err
	}
	return new(empty.Empty), nil
}

func (handler *UserHandler) UpdateUser(ctx context.Context, user *models.UserUpdate) (*empty.Empty, error) {
	_, err := handler.userService.UpdateUser(user)

	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func (handler *UserHandler) DeleteUser(ctx context.Context, userId *models.UserId) (*empty.Empty, error) {
	err := handler.userService.DeleteUser(userId)

	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func mustEmbedUnimplementedUsersServer() {
}
