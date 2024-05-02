package entities

import "github.com/google/uuid"

type GetTeacherIdByUserIdRequest struct {
	UserId uuid.UUID `json:"user_id"`
}

type GetCourseByTeacherIdRequest struct {
	UserId uuid.UUID `json:"user_id"`
}

type GetCourseByCourseCodeRequest struct {
	CourseCode string `json:"course_code"`
}

type CreateCourseRequest struct {
	UserId uuid.UUID `json:"user_id"`
}
