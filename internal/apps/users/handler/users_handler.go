package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/users"
	"github.com/tnnz20/godemy-be/internal/apps/users/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/response"
)

type handler struct {
	svc users.Service
}

func NewHandler(svc users.Service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) Register(c *fiber.Ctx) error {
	var req entities.RegisterPayload

	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	if err := h.svc.Register(c.UserContext(), req); err != nil {
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

func (h handler) Login(c *fiber.Ctx) error {
	var req entities.LoginPayload

	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBadRequest(c, err)
	}

	token, err := h.svc.Login(c.UserContext(), req)
	if err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 400:
			return response.ErrorBadRequest(c, err)
		case 401:
			return response.ErrorUnauthorized(c)
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, token)
}

func (h handler) GetUser(c *fiber.Ctx) error {
	id := c.Locals("id").(string)

	userId, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorBadRequest(c, err)
	}

	var req entities.GetUserPayload
	req.ID = userId

	user, err := h.svc.GetUser(c.UserContext(), req)
	if err != nil {
		errorMapping := errs.ErrorMapping[err]
		switch errorMapping {
		case 404:
			return response.ErrorNotFound(c, err)
		default:
			return response.InternalServerError(c, err)
		}
	}

	return response.SuccessOK(c, user)
}
