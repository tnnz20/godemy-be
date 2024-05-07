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

	course := entities.Courses{
		UsersId:    req.UsersId,
		CourseCode: courseCode,
		CourseName: "golang-fundamental",
	}

	if err := course.Validate(); err != nil {
		return err
	}

	if err := s.Repository.CreateCourse(ctx, course); err != nil {
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
		if err == sql.ErrNoRows {
			err = errs.ErrCourseNotFound
			return
		}
		return
	}

	res = entities.CourseResponse(course)

	return
}

func (s *service) GetCoursesByUsersIdWithPagination(ctx context.Context, req entities.GetCoursesByUsersIdWithPaginationPayload) (res []entities.CourseResponse, err error) {
	coursePagination := entities.CoursesPagination{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	courses, err := s.Repository.FindCoursesByUsersIdWithPagination(ctx, req.UsersId, coursePagination)
	if err != nil {
		return
	}

	fmt.Println(courses)

	if len(courses) == 0 {
		err = errs.ErrCourseEmpty
		return
	}

	for _, course := range courses {
		res = append(res, entities.CourseResponse(course))
	}

	return
}
