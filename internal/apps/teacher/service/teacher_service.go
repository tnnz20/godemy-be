package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/tnnz20/godemy-be/internal/apps/teacher"
	"github.com/tnnz20/godemy-be/internal/apps/teacher/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/helpers"
)

type service struct {
	teacher.Repository
}

func NewService(teacherRepo teacher.Repository) teacher.Service {
	return &service{
		Repository: teacherRepo,
	}
}

func (s *service) CreateCourse(ctx context.Context, req entities.CreateCourseRequest) (err error) {
	teacher, err := s.Repository.FindTeacherIdByUserId(ctx, req.UserId)
	if err != nil {
		return
	}

	// Generate course code
	randomString := helpers.GenerateRandomString(7)
	courseCode := fmt.Sprintf("go-%s", randomString)

	// TODO: check if course code exist, regenerate

	course := entities.Course{
		CourseName: "golang-fundamental",
		CourseCode: courseCode,
		TeacherId:  teacher.ID,
	}

	if err := course.Validate(); err != nil {
		return err
	}

	newCourse := s.Repository.CreateCourse(ctx, course)
	return newCourse
}

func (s *service) GetTeacherIdByUserId(ctx context.Context, req entities.GetTeacherIdByUserIdRequest) (res entities.GetTeacherIdByUserIdResponse, err error) {
	teacher, err := s.Repository.FindTeacherIdByUserId(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrTeacherNotFound
			return entities.GetTeacherIdByUserIdResponse{}, err
		}
		return entities.GetTeacherIdByUserIdResponse{}, err
	}

	res = entities.GetTeacherIdByUserIdResponse{
		Id:     teacher.ID,
		UserId: req.UserId,
	}

	return
}
