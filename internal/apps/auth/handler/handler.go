package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnnz20/godemy-be/internal/apps/auth"
	"github.com/tnnz20/godemy-be/internal/apps/auth/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/response"
)

type Handler struct {
	svc auth.Service
}

func NewHandler(svc auth.Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) Register(c *fiber.Ctx) error {
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

	return response.SuccessCreated(c, nil)
}

func (h Handler) Login(c *fiber.Ctx) error {
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
