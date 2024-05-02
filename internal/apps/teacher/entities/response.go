package entities

import (
	"time"

	"github.com/google/uuid"
)

type GetTeacherIdByUserIdResponse struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
}

type GetCourseByTeacherIdResponse struct {
	ID         uuid.UUID `json:"id"`
	TeacherId  uuid.UUID `json:"teacher_id"`
	CourseName string    `json:"course_name"`
	CourseCode string    `json:"course_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetCourseByCourseCodeResponse struct {
	ID         uuid.UUID `json:"id"`
	TeacherId  uuid.UUID `json:"teacher_id"`
	CourseName string    `json:"course_name"`
	CourseCode string    `json:"course_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
