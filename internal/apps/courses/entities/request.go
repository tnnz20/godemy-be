package entities

import "github.com/google/uuid"

type CreateCoursePayload struct {
	UsersId uuid.UUID `json:"users_id"`
	// CourseName string    `json:"course_name"`
}

type GetCourseByCourseCodePayload struct {
	CourseCode string `json:"course_code"`
}

type GetCoursesByUsersIdWithPaginationPayload struct {
	UsersId uuid.UUID `json:"users_id"`
	Limit   int       `query:"limit" json:"limit"`
	Offset  int       `query:"offset" json:"offset"`
}

func (p *GetCoursesByUsersIdWithPaginationPayload) GenerateDefaultValue() GetCoursesByUsersIdWithPaginationPayload {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	return *p
}
