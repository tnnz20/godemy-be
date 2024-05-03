package entities

import "github.com/google/uuid"

type GetTeacherIdByUserIdPayload struct {
	UserId uuid.UUID `json:"user_id"`
}

type GetCourseByTeacherIdPayload struct {
	UserId uuid.UUID `json:"user_id"`
}

type GetCourseByCourseCodePayload struct {
	CourseCode string `json:"course_code"`
}

type CreateCoursePayload struct {
	UserId uuid.UUID `json:"user_id"`
}
