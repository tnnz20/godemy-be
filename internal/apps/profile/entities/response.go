package entities

import (
	"time"
)

type GetProfileByUserIdResponse struct {
	Name       string
	Date       time.Time
	Address    string
	Gender     string
	ProfileImg string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
