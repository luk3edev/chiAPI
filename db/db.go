package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Context *gorm.DB
}

const ( // for now, TODO: move to configs
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "myappdb"
)

func NewDatabase() Database {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database..")
	}
	return Database{Context: db}
}

func NewTestDatabase() Database {
	// Set up the test database
	db, err := gorm.Open(sqlite.Open(":memory"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database..")
	}

	return Database{Context: db}
}
