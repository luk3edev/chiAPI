package main

import (
	"chiAPI/db"
	"chiAPI/handlers"
	"chiAPI/services"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"testing"
)

func main() {
	// Initialize database singleton
	fmt.Println("Initializing database...")

	var database = SetupDatabase()

	// Initialize user service with database dependency
	userService := services.NewUserService(&database)

	// Initialize user handler with service dependency
	r := chi.NewRouter()

	// Middleware setup
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initialize user handler with service dependency
	userHandler := handlers.NewUserHandler(userService)

	// User routes
	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Delete("/{id}", userHandler.DeleteUser)
		r.Put("/{id}", userHandler.UpdateUser)
	})

	// Start server
	fmt.Println("🚀 Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("failed to start server: " + err.Error())
	}
}

func SetupDatabase() db.Database {
	var database db.Database
	if isTest() {
		database = db.NewTestDatabase()
		db.Migrate(database)
	} else {
		database = db.NewDatabase()
		db.Migrate(database)
		db.SetupUser(database)
	}
	return database
}

func isTest() bool {
	return os.Getenv("GO_ENV") == "test" || testing.Short()
}
