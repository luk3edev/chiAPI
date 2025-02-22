package services

import (
	"chiAPI/db"
	"chiAPI/models"
	"fmt"
	"github.com/google/uuid"
)

type IUserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id string) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
	DeleteUser(id string) error
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

func (s *UserService) GetUser(id string) (models.User, error) {
	var user models.User

	parsedId, err := uuid.Parse(id)

	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return models.User{}, err
	}

	if err := s.db.Context.First(&user, "id = ?", parsedId).Error; err != nil {
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

	result, err := s.GetUser(id)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

func (s *UserService) DeleteUser(id string) error {
	if err := s.db.Context.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
