package main

import (
	"log"
	"net/http"
	"os"
	_ "user-service/docs" // This is required for swagger
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"
	"user-service/pkg/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           User Service API
// @version         1.0
// @description     A User management service API in Go.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	dbConfig := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Setup router
	router := mux.NewRouter()

	// Swagger documentation
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// API routes
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", userHandler.ListUsers).Methods("GET")

	// Start server
	log.Printf("Server starting on port 8080")
	log.Printf("Swagger documentation available at http://localhost:8080/swagger/index.html")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
