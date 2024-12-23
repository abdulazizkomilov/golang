package domain

import "time"

// User represents the user model
type User struct {
	ID        uint      `json:"id" example:"1"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"john.doe@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	FirstName string `json:"first_name" example:"John" validate:"required"`
	LastName  string `json:"last_name" example:"Doe" validate:"required"`
	Email     string `json:"email" example:"john.doe@example.com" validate:"required,email"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"john.doe@example.com" validate:"email"`
}
