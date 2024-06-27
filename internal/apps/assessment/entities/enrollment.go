package entities

import (
	"github.com/google/uuid"
)

type Enrollment struct {
	ID        uuid.UUID
	UsersId   uuid.UUID
	CoursesId uuid.UUID
	Progress  uint8
	CreatedAt int64
	UpdatedAt int64
}
