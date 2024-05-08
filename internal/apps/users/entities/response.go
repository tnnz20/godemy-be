package entities

import (
	"time"

	"github.com/google/uuid"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Date       time.Time `json:"date"`
	Address    string    `json:"address"`
	Gender     string    `json:"gender"`
	ProfileImg string    `json:"profile_img"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
