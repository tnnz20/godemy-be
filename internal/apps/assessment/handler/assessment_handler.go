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

	var req entities.CreateAssessmentRequest
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

func (h handler) GetAssessmentsFiltered(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetAssessmentRequest
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

func (h handler) GetAssessmentByAssessmentCode(c *fiber.Ctx) error {
	// Get the user id from the context
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetAssessmentByAssessmentCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	req.UsersId = userId

	res, err := h.Service.GetAssessmentResultByAssessmentCode(c.UserContext(), req)
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