package entities

import (
	"time"

	"github.com/google/uuid"
)

type Enrollment struct {
	ID        uuid.UUID
	UsersId   uuid.UUID
	CoursesId uuid.UUID
	Progress  uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
