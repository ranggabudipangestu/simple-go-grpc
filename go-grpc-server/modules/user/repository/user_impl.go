package repository

import (
	"github.com/ranggabudipangestu/go-grpc-server/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) AddUser(user *models.User) (*models.User, error) {
	dataDB := models.UserDB{
		Name:     user.Name,
		Email:    user.Email,
		Alamat:   user.Alamat,
		Password: user.Password,
	}
	err := r.db.Table("user").Save(&dataDB).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindUserById(userId *models.UserId) (*models.User, error) {
	var user models.User
	err := r.db.Table("user").Where("id = ?", userId.Id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindUsers() (*[]models.User, error) {
	var users []models.User

	err := r.db.Table("user").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}
func (r *UserRepositoryImpl) UpdateUser(user *models.UserUpdate) (*models.User, error) {
	var dataUser models.User

	err := r.db.Table("user").Where("id=?", user.Id).First(&dataUser).Updates(&user.User).Error

	if err != nil {
		return nil, err
	}

	return &dataUser, nil

}

func (r *UserRepositoryImpl) DeleteUser(userId *models.UserId) error {
	var user models.User

	err := r.db.Table("user").Where("id=?", userId.Id).First(&user).Delete(&user).Error

	if err != nil {
		return err
	}

	return nil

}
