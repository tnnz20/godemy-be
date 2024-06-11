package entities

import "github.com/google/uuid"

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

type CreateCoursePayload struct {
	UsersId    uuid.UUID `json:"users_id"`
	CourseName string    `json:"course_name"`
}

type GetCourseByCourseCodePayload struct {
	CourseCode string `json:"course_code"`
}

type GetCoursesByUsersIdPayload struct {
	UsersId    uuid.UUID `json:"users_id"`
	CourseName string    `query:"course_name"`
}

type GetEnrolledUsersByUsersIdPayload struct {
	UsersId uuid.UUID `json:"users_id"`
}

type GetCoursesByUsersIdWithPaginationPayload struct {
	UsersId    uuid.UUID `json:"users_id"`
	CourseName string    `query:"course_name"`
	ModelPaginationPayload
}

type GetTotalCoursesByUsersIdPayload struct {
	UsersId    uuid.UUID `json:"users_id"`
	CourseName string    `query:"course_name"`
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

type GetTotalEnrolledUsersByCourseIdPayload struct {
	CourseId uuid.UUID `params:"courseId"`
	Name     string    `query:"name"`
}
