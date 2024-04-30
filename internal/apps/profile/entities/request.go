package entities

import (
	"time"

	"github.com/google/uuid"
)

type GetProfileByUserIdRequest struct {
	UserId uuid.UUID `json:"user_id"`
}

type UpdateProfileRequest struct {
	Name       string    `json:"name"`
	Date       time.Time `json:"date"`
	Address    string    `json:"address"`
	Gender     string    `json:"gender"`
	ProfileImg string    `json:"profile_img"`
	UserID     uuid.UUID `json:"omitempty"`
}
