package auth

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tnnz20/godemy-be/util"
)

type Handler struct {
	AuthService Service
	Validate    *validator.Validate
}

func NewHandler(s Service, validate *validator.Validate) *Handler {
	return &Handler{
		AuthService: s,
		Validate:    validate,
	}
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.AuthService.SignIn(c.Context(), &req)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "User not found.")
		} else if err.Error() == "Invalid Password" {
			return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
		}

		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusOK, "Successfully Sign-in.", res)
}
