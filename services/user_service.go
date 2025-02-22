package services

import (
	"chiAPI/db"
	"chiAPI/models"
)

type IUserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
}

type UserService struct {
	db *db.Database
}

func NewUserService(db *db.Database) IUserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	if err := s.db.Context.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUser(id int) (models.User, error) {
	var user models.User
	if err := s.db.Context.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Context.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(id string, user models.User) (models.User, error) {
	if err := s.db.Context.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
