package entities

import (
	"time"

	"github.com/google/uuid"
)

type AssessmentResponse struct {
	ID              uuid.UUID `json:"id"`
	UsersId         uuid.UUID `json:"users_id"`
	CoursesId       uuid.UUID `json:"courses_id"`
	AssessmentValue float32   `json:"assessment_value"`
	AssessmentCode  string    `json:"assessment_code"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
