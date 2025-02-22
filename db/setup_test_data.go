package db

import (
	"chiAPI/models"
	"github.com/google/uuid"
)

func SetupUser(database Database) {
	user_id, _ := uuid.NewRandom()

	user := &models.User{
		ID:        user_id,
		FirstName: "testName",
		LastName:  "testLastName",
	}
	if err := database.Context.Create(user); err != nil {
		panic(err)
	}
}
