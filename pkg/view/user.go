package view

import "time"

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
