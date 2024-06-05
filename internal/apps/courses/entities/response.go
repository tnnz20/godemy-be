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

type CoursesLengthResponse struct {
	Total int `json:"total"`
}

type EnrolledUsersResponse struct {
	ID         uuid.UUID `json:"id"`
	CourseName string    `json:"course_name"`
	Name       string    `json:"name"`
	Progress   uint8     `json:"progress"`
	UpdateAt   time.Time `json:"updated_at"`
}

type EnrolledUsersLengthResponse struct {
	Total int `json:"total"`
}
