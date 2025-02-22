package db

import (
	"chiAPI/models"
	"fmt"
)

func Migrate(database Database) {
	if err := database.Context.AutoMigrate(&models.User{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migrated")
}
