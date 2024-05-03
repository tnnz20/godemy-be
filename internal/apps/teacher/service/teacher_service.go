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

func (s *service) GetCourseByTeacherId(ctx context.Context, req entities.GetCourseByTeacherIdPayload) (res entities.GetCourseByTeacherIdResponse, err error) {
	teacher, err := s.Repository.FindTeacherIdByUserId(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrTeacherNotFound
			return entities.GetCourseByTeacherIdResponse{}, err
		}
		return entities.GetCourseByTeacherIdResponse{}, err
	}

	NewCourse, err := s.Repository.FindCourseByTeacherId(ctx, teacher.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errs.ErrCourseNotFound
			return
		}
		return
	}

	res = entities.GetCourseByTeacherIdResponse(NewCourse)
	return
}

func (s *service) GetCourseByCourseCode(ctx context.Context, req entities.GetCourseByCourseCodePayload) (res entities.GetCourseByCourseCodeResponse, err error) {
	course := entities.Course{
		CourseCode: req.CourseCode,
	}

	if err := course.ValidateCourseCode(); err != nil {
		return entities.GetCourseByCourseCodeResponse{}, err
	}

	NewCourses, err := s.Repository.FindCourseByCourseCode(ctx, course.CourseCode)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errs.ErrCourseNotFound
			return
		}
		return
	}

	res = entities.GetCourseByCourseCodeResponse(NewCourses)

	return
}

func (s *service) CreateCourse(ctx context.Context, req entities.CreateCoursePayload) (err error) {
	teacher, err := s.Repository.FindTeacherIdByUserId(ctx, req.UserId)
	if err != nil {
		return
	}

	var courseCode string
	for {
		randomString := helpers.GenerateRandomString(7)
		courseCode = fmt.Sprintf("go-%s", randomString)

		checkCourse, err := s.Repository.FindCourseByCourseCode(ctx, courseCode)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
		}

		if !checkCourse.IsCourseCodeExist() {
			// Course code doesn't exist, so it's unique
			break
		}
		// Course code already exists, so try again
	}

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

func (s *service) GetTeacherIdByUserId(ctx context.Context, req entities.GetTeacherIdByUserIdPayload) (res entities.GetTeacherIdByUserIdResponse, err error) {
	teacher, err := s.Repository.FindTeacherIdByUserId(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrTeacherNotFound
			return entities.GetTeacherIdByUserIdResponse{}, err
		}
		return entities.GetTeacherIdByUserIdResponse{}, err
	}

	res = entities.GetTeacherIdByUserIdResponse{
		ID:     teacher.ID,
		UserId: req.UserId,
	}

	return
}
