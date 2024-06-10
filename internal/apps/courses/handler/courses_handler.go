package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/courses"
	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/response"
)

type handler struct {
	courses.Service
}

func NewHandler(svc courses.Service) handler {
	return handler{
		Service: svc,
	}
}

func (h handler) CreateCourse(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.CreateCoursePayload
	req.UsersId = userId

	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	if err := h.Service.CreateCourse(c.UserContext(), req); err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 409:
			return response.ErrorConflict(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessCreated(c)
}

func (h handler) GetCoursesByUsersId(c *fiber.Ctx) error {
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}
	limit := c.QueryInt("limit", 6)
	offset := c.QueryInt("offset", 0)

	var req entities.GetCoursesByUsersIdWithPaginationPayload
	req.CourseName = c.Query("course_name")
	req.UsersId = userId
	req.Limit = limit
	req.Offset = offset

	courses, err := h.Service.GetCoursesByUsersIdWithPagination(c.UserContext(), req)
	if err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, courses)
}

func (h handler) GetTotalCourses(c *fiber.Ctx) error {
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetTotalCoursesByUsersIdPayload
	req.UsersId = userId
	req.CourseName = c.Query("course_name")

	total, err := h.Service.GetTotalCourses(c.UserContext(), req)
	if err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, total)
}

func (h handler) EnrollCourse(c *fiber.Ctx) error {
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.EnrollCoursePayload
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.UsersId = userId

	if err := h.Service.EnrollCourse(c.UserContext(), req); err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		case 409:
			return response.ErrorConflict(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessCreated(c)
}

func (h handler) GetCourseEnrollmentDetail(c *fiber.Ctx) error {
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetCourseEnrollmentByUsersIdPayload
	req.UsersId = userId

	enrollment, err := h.Service.GetCourseEnrollmentByUsersId(c.UserContext(), req)
	if err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, enrollment)
}

func (h handler) UpdateProgressCourseEnrollment(c *fiber.Ctx) error {
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.UpdateEnrollmentProgressPayload
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.UsersId = userId

	if err := h.Service.UpdateProgressCourseEnrollment(c.UserContext(), req); err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, nil)
}

func (h handler) GetEnrolledUsers(c *fiber.Ctx) error {

	var req entities.GetEnrolledUsersByCourseIdPayload
	limit := c.QueryInt("limit", 6)
	offset := c.QueryInt("offset", 0)

	paramsCourseId := c.Params("courseId")

	courseId, err := uuid.Parse(paramsCourseId)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.Name = c.Query("name")
	req.CourseId = courseId
	req.Limit = limit
	req.Offset = offset

	users, err := h.Service.GetEnrolledUsersByCourseId(c.UserContext(), req)
	if err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, users)
}

func (h handler) GetTotalEnrolledUsers(c *fiber.Ctx) error {

	var req entities.GetTotalEnrolledUsersByCourseIdPayload

	paramsCourseId := c.Params("courseId")

	courseId, err := uuid.Parse(paramsCourseId)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.Name = c.Query("name")
	req.CourseId = courseId

	total, err := h.Service.GetTotalEnrolledUsersByCourseId(c.UserContext(), req)
	if err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, total)
}
