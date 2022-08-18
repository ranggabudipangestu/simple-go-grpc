package repository

import "github.com/ranggabudipangestu/go-grpc-server/models"

type UserRepository interface {
	AddUser(user *models.User) (*models.User, error)
	FindUserById(id *models.UserId) (*models.User, error)
	FindUsers() (*[]models.User, error)
	UpdateUser(user *models.UserUpdate) (*models.User, error)
	DeleteUser(userId *models.UserId) error
}
