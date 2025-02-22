package db

import (
	"chiAPI/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetupUser(context *gorm.DB) {
	uuid, _ := uuid.NewRandom()

	user := &models.User{
		ID:        uuid,
		FirstName: "testName",
		LastName:  "testLastName",
	}
	if err := context.Create(user); err != nil {
		panic(err)
	}
}
