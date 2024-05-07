package entities

import (
	"time"

	"github.com/google/uuid"
)

type CourseResponse struct {
	ID         uuid.UUID `json:"id"`
	UsersId    uuid.UUID `json:"users_id"`
	CourseName string    `json:"course_name"`
	CourseCode string    `json:"course_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}