package entities

import "github.com/google/uuid"

type CreateCoursePayload struct {
	UsersId    uuid.UUID `json:"users_id"`
	CourseName string    `json:"course_name"`
}

type GetCourseByCourseCodePayload struct {
	CourseCode string `json:"course_code"`
}

type GetCoursesByUsersIdPayload struct {
	UsersId uuid.UUID `json:"users_id"`
}

type GetEnrolledUsersByUsersIdPayload struct {
	UsersId uuid.UUID `json:"users_id"`
}

type ModelPaginationPayload struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type GetCoursesByUsersIdWithPaginationPayload struct {
	UsersId uuid.UUID `json:"users_id"`
	ModelPaginationPayload
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

type EnrollCoursePayload struct {
	UsersId    uuid.UUID `json:"users_id"`
	CourseCode string    `json:"course_code"`
}

type UpdateEnrollmentProgressPayload struct {
	UsersId  uuid.UUID `json:"users_id"`
	Progress uint8     `json:"progress"`
}

type GetCourseEnrollmentByUsersIdPayload struct {
	UsersId uuid.UUID `json:"users_id"`
}

type GetEnrolledUsersByCourseIdPayload struct {
	CourseId uuid.UUID `params:"courseId"`
	Name     string    `query:"name"`
	ModelPaginationPayload
}
