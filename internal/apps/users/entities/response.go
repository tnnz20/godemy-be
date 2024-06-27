package entities

import (
	"github.com/google/uuid"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Date       int64     `json:"date"`
	Address    string    `json:"address"`
	Gender     string    `json:"gender"`
	ProfileImg string    `json:"profile_img"`
	CreatedAt  int64     `json:"created_at"`
	UpdatedAt  int64     `json:"updated_at"`
}
