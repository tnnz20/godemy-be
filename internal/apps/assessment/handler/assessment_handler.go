package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/assessment"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/response"
)

type handler struct {
	assessment.Service
}

func NewHandler(svc assessment.Service) handler {
	return handler{
		Service: svc,
	}
}

func (h handler) CreateAssessment(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.CreateAssessmentPayload
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.UsersId = userId

	if err := h.Service.CreateAssessmentResult(c.UserContext(), req); err != nil {
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

func (h handler) GetAssessments(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetAssessmentPayload
	req.UsersId = userId

	res, err := h.Service.GetAssessmentsResult(c.UserContext(), req)
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

	return response.SuccessOK(c, res)
}

func (h handler) GetFilteredAssessmentResult(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetAssessmentResultWithPaginationPayload

	req.AssessmentCode = c.Query("assessment_code")
	req.Limit = c.QueryInt("limit", 5)
	req.Offset = c.QueryInt("offset", 0)
	req.UsersId = userId

	res, err := h.Service.GetFilteredAssessmentResult(c.UserContext(), req)
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

	return response.SuccessOK(c, res)
}

func (h handler) GetTotalFilteredAssessmentResult(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetAssessmentResultWithPaginationPayload
	req.UsersId = userId
	req.AssessmentCode = c.Query("assessment_code")

	res, err := h.Service.GetTotalFilteredAssessmentResult(c.UserContext(), req)
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

	return response.SuccessOK(c, res)
}

func (h handler) GetAssessmentsResultUsers(c *fiber.Ctx) error {
	var req entities.GetAssessmentResultsByCourseIdPayload

	paramsCourseId := c.Params("courseId")
	if paramsCourseId == "" {
		err := errs.ErrCourseIdRequired
		return response.ErrorBadRequest(c, err)
	}

	courseId, err := uuid.Parse(paramsCourseId)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.CoursesId = courseId
	req.AssessmentCode = c.Query("assessment_code")
	req.Name = c.Query("name")
	req.Status = uint8(c.QueryInt("status"))
	req.Sort = c.Query("sort", "ASC")
	req.Limit = c.QueryInt("limit", 6)
	req.Offset = c.QueryInt("offset", 0)

	res, err := h.Service.GetAssessmentsResultUsers(c.UserContext(), req)
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

	return response.SuccessOK(c, res)
}

func (h handler) GetTotalAssessmentsResultUsers(c *fiber.Ctx) error {
	var req entities.GetAssessmentResultsByCourseIdPayload

	paramsCourseId := c.Params("courseId")
	if paramsCourseId == "" {
		err := errs.ErrCourseIdRequired
		return response.ErrorBadRequest(c, err)
	}

	courseId, err := uuid.Parse(paramsCourseId)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.CoursesId = courseId
	req.AssessmentCode = c.Query("assessment_code")
	req.Name = c.Query("name")

	res, err := h.Service.GetTotalAssessmentsResultUsers(c.UserContext(), req)
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

	return response.SuccessOK(c, res)
}

func (h handler) CreateUsersAssessment(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.CreateUsersAssessmentPayload
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.UsersId = userId
	if err := h.Service.CreateUsersAssessment(c.UserContext(), req); err != nil {
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

func (h handler) UpdateUsersAssessmentStatus(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.UpdateUsersAssessmentStatusPayload
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.UsersId = userId

	if err := h.Service.UpdateUsersAssessmentStatus(c.UserContext(), req); err != nil {
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

func (h handler) GetUsersAssessment(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetUsersAssessmentPayload
	if err := c.QueryParser(&req); err != nil {
		err = errs.ErrAssessmentCodeRequired
		return response.ErrorBadRequest(c, err)
	}

	req.UsersId = userId

	res, err := h.Service.GetUsersAssessment(c.UserContext(), req)
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

	return response.SuccessOK(c, res)
}
