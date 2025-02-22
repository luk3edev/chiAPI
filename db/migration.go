package db

import (
	"chiAPI/models"
	"fmt"
	"gorm.io/gorm"
)

func Migrate(context *gorm.DB) {
	if err := context.AutoMigrate(&models.User{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migrated")
}
