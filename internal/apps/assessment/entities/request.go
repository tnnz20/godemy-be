package entities

import "github.com/google/uuid"

type CreateAssessmentPayload struct {
	UsersId         uuid.UUID `json:"users_id"`
	AssessmentValue float32   `json:"assessment_value"`
	AssessmentCode  string    `json:"assessment_code"`
}

type GetAssessmentPayload struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `query:"assessment_code"`
}

type GetAssessmentResultWithPaginationPayload struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `query:"assessment_code"`
	ModelPaginationPayload
}

type GetAssessmentResultsByCourseIdPayload struct {
	CoursesId      uuid.UUID `json:"courses_id"`
	Name           string    `query:"name"`
	AssessmentCode string    `query:"assessment_code"`
	Status         uint8     `query:"status"`
	ModelPaginationPayload
}

type CreateUsersAssessmentPayload struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
	RandomArrayId  []uint8   `json:"random_array_id"`
}

type GetUsersAssessmentPayload struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `query:"assessment_code"`
}

type UpdateUsersAssessmentStatusPayload struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
	Status         uint8     `json:"status"`
}

type ModelPaginationPayload struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func (p *ModelPaginationPayload) GenerateDefaultValue() ModelPaginationPayload {
	// limit rows
	if p.Limit <= 0 {
		p.Limit = 10
	}

	// skip rows
	if p.Offset <= 0 {
		p.Offset = 0
	}

	return *p
}
