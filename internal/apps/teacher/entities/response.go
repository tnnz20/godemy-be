package entities

import (
	"time"

	"github.com/google/uuid"
)

type GetTeacherIdByUserIdResponse struct {
	Id     uuid.UUID
	UserId uuid.UUID
}

type GetCourseByTeacherIdResponse struct {
	Id         uuid.UUID
	TeacherId  uuid.UUID
	CodeCourse string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
