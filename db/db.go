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

func NewDbConnection() Database {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database..")
	}
	return Database{Context: db}
}

func SetupTestDatabase() (*gorm.DB, error) {
	// Setup the test database
	db, err := gorm.Open(sqlite.Open(":memory"), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to the database..")
		return nil, err
	}

	Migrate(db)

	return db, nil
}
