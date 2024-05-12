package entities

import "github.com/google/uuid"

type CreateAssessmentRequest struct {
	UserId          uuid.UUID `json:"users_id"`
	AssessmentValue float32   `json:"assessment_value"`
	AssessmentCode  string    `json:"assessment_code"`
}

type GetAssessmentRequest struct {
	UserId uuid.UUID `json:"users_id"`
}

type GetAssessmentByAssessmentCodeRequest struct {
	UserId         uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
}
