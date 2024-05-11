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

type CourseEnrollmentResponse struct {
	ID        uuid.UUID `json:"id"`
	UsersId   uuid.UUID `json:"users_id"`
	CoursesId uuid.UUID `json:"courses_id"`
	Progress  uint8     `json:"progress"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListUserCourseEnrollmentResponse struct {
	ID       uuid.UUID `json:"id"`
	CourseId string    `json:"course_id"`
	Name     string    `json:"name"`
	Progress uint8     `json:"progress"`
}
