package entities

import "github.com/google/uuid"

type GetTeacherIdByUserIdRequest struct {
	UserId uuid.UUID
}

type GetCourseByTeacherId struct {
	UserId uuid.UUID
}

type CreateCourseRequest struct {
	UserId uuid.UUID
}
