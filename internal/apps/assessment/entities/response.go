package entities

import (
	"github.com/google/uuid"
)

type AssessmentResponse struct {
	ID              uuid.UUID `json:"id"`
	UsersId         uuid.UUID `json:"users_id"`
	CoursesId       uuid.UUID `json:"courses_id"`
	AssessmentValue float32   `json:"assessment_value"`
	AssessmentCode  string    `json:"assessment_code"`
	Status          uint8     `json:"status"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
}

type AssessmentUserResponse struct {
	ID             uuid.UUID `json:"id"`
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
	RandomArrayId  []uint8   `json:"random_array_id"`
	Status         string    `json:"status"`
	CreatedAt      int64     `json:"created_at"`
	UpdatedAt      int64     `json:"updated_at"`
}

type AssessmentTotalResponse struct {
	Total int `json:"total"`
}

type AssessmentResultUsersResponse struct {
	ID              uuid.UUID `json:"id"`
	UsersId         uuid.UUID `json:"users_id"`
	Name            string    `json:"name"`
	CoursesId       uuid.UUID `json:"courses_id"`
	AssessmentValue float32   `json:"assessment_value"`
	AssessmentCode  string    `json:"assessment_code"`
	Status          uint8     `json:"status"`
	CreatedAt       int64     `json:"created_at"`
}
