package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/tnnz20/godemy-be/internal/apps/courses"
	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/helpers"
)

type service struct {
	courses.Repository
}

func NewService(coursesRepo courses.Repository) courses.Service {
	return &service{
		Repository: coursesRepo,
	}
}

func (s *service) CreateCourse(ctx context.Context, req entities.CreateCoursePayload) (err error) {
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
	}

	NewCourse := entities.NewCourses(req.UsersId, req.CourseName, courseCode)

	if err := NewCourse.Validate(); err != nil {
		return err
	}

	if err := s.Repository.CreateCourse(ctx, NewCourse); err != nil {
		return err
	}

	return
}

func (s *service) GetCourseByCourseCode(ctx context.Context, req entities.GetCourseByCourseCodePayload) (res entities.CourseResponse, err error) {
	courseEntity := entities.Courses{
		CourseCode: req.CourseCode,
	}

	if err := courseEntity.ValidateCourseCode(); err != nil {
		return entities.CourseResponse{}, err
	}

	course, err := s.Repository.FindCourseByCourseCode(ctx, courseEntity.CourseCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrCourseNotFound
			return
		}
		return
	}

	res = entities.CourseResponse(course)

	return
}

func (s *service) GetCoursesByUsersIdWithPagination(ctx context.Context, req entities.GetCoursesByUsersIdWithPaginationPayload) (res []entities.CourseResponse, err error) {
	NewCoursePagination := entities.NewCoursesPagination(req.Limit, req.Offset)

	courses, err := s.Repository.FindCoursesByUsersIdWithPagination(ctx, req.UsersId, req.CourseName, NewCoursePagination)
	if err != nil {
		return
	}

	if len(courses) == 0 {
		err = errs.ErrCourseEmpty
		return
	}

	for _, course := range courses {
		res = append(res, entities.CourseResponse(course))
	}

	return
}

func (s *service) GetCoursesByUsersId(ctx context.Context, req entities.GetCoursesByUsersIdPayload) (res []entities.CourseResponse, err error) {

	courses, err := s.Repository.FindCoursesByUsersId(ctx, req.UsersId, req.CourseName)
	if err != nil {
		return
	}

	if len(courses) == 0 {
		err = errs.ErrCourseEmpty
		return
	}

	for _, course := range courses {
		res = append(res, entities.CourseResponse(course))
	}

	return
}

func (s *service) GetTotalCourses(ctx context.Context, req entities.GetTotalCoursesByUsersIdPayload) (res entities.CoursesLengthResponse, err error) {
	total, err := s.Repository.FindTotalCoursesByUsersId(ctx, req.UsersId, req.CourseName)
	if err != nil {
		return
	}

	if total == 0 {
		err = errs.ErrCourseEmpty
		return
	}

	res = entities.CoursesLengthResponse{
		Total: total,
	}

	return
}

func (s *service) EnrollCourse(ctx context.Context, req entities.EnrollCoursePayload) (err error) {

	validateCourse := entities.Courses{
		CourseCode: req.CourseCode,
	}

	if err := validateCourse.ValidateCourseCode(); err != nil {
		return err
	}

	userEnroll, err := s.Repository.FindCourseEnrollmentByUsersId(ctx, req.UsersId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if userEnroll.IsEnrolled() {
		err = errs.ErrUserAlreadyEnrolled
		return
	}

	course, err := s.Repository.FindCourseByCourseCode(ctx, validateCourse.CourseCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrCourseNotFound
			return
		}
		return
	}

	NewEnrollment := entities.NewEnrollment(req.UsersId, course.ID)

	if err := s.Repository.InsertCourseEnrollment(ctx, NewEnrollment); err != nil {
		return err
	}

	return
}

func (s *service) GetCourseEnrollmentByUsersId(ctx context.Context, req entities.GetCourseEnrollmentByUsersIdPayload) (res entities.CourseEnrollmentResponse, err error) {

	enrollment, err := s.Repository.FindCourseEnrollmentByUsersId(ctx, req.UsersId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrCourseEnrollmentNotFound
			return
		}
		return
	}

	res = entities.CourseEnrollmentResponse(enrollment)

	return
}

func (s *service) UpdateProgressCourseEnrollment(ctx context.Context, req entities.UpdateEnrollmentProgressPayload) (err error) {
	enrollment, err := s.Repository.FindCourseEnrollmentByUsersId(ctx, req.UsersId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrCourseEnrollmentNotFound
			return
		}
		return
	}

	if err := enrollment.UpdateProgress(req.Progress); err != nil {
		return err
	}

	if err := s.Repository.UpdateEnrollmentProgress(ctx, enrollment); err != nil {
		return err
	}

	return
}

func (s *service) GetEnrolledUsersByCourseId(ctx context.Context, req entities.GetEnrolledUsersByCourseIdPayload) (res []entities.EnrolledUsersResponse, err error) {
	NewCoursePagination := entities.NewCoursesPagination(req.Limit, req.Offset)

	courses, err := s.Repository.FindEnrolledUsersByCourseId(ctx, req.CourseId, req.Name, NewCoursePagination)
	if err != nil {
		return
	}

	if len(courses) == 0 {
		err = errs.ErrCourseEmpty
		return
	}

	res = append(res, courses...)

	return
}

func (s *service) GetTotalEnrolledUsersByCourseId(ctx context.Context, req entities.GetTotalEnrolledUsersByCourseIdPayload) (res entities.EnrolledUsersLengthResponse, err error) {
	total, err := s.Repository.FindTotalEnrolledUsersByCourseId(ctx, req.CourseId, req.Name)
	if err != nil {
		return
	}

	if total == 0 {
		err = errs.ErrCourseEmpty
		return
	}

	res = entities.EnrolledUsersLengthResponse{
		Total: total,
	}

	return
}
