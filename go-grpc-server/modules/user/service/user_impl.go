package service

import (
	"github.com/ranggabudipangestu/go-grpc-server/models"
	"github.com/ranggabudipangestu/go-grpc-server/modules/user/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
	}
}

func (s *UserServiceImpl) AddUser(user *models.User) (*models.User, error) {
	return s.UserRepository.AddUser(user)
}
func (s *UserServiceImpl) FindUserById(id *models.UserId) (*models.User, error) {
	return s.UserRepository.FindUserById(id)
}
func (s *UserServiceImpl) FindUsers() (*[]models.User, error) {
	return s.UserRepository.FindUsers()
}
func (s *UserServiceImpl) UpdateUser(user *models.UserUpdate) (*models.User, error) {
	return s.UserRepository.UpdateUser(user)
}
func (s *UserServiceImpl) DeleteUser(userId *models.UserId) error {
	return s.UserRepository.DeleteUser(userId)
}
