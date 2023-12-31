package student

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tnnz20/godemy-be/util"
)

type Handler struct {
	StudentService Service
	Validate       *validator.Validate
}

func NewHandler(s Service, validate *validator.Validate) *Handler {
	return &Handler{
		StudentService: s,
		Validate:       validate,
	}
}

const (
	unauthorizedMsg = "Unauthorized not a "
	notFoundMsg     = " not found."
)

func (h *Handler) GetStudentByUserId(c *fiber.Ctx) error {

	id, role := util.JwtClaim(c)

	if role != "student" {
		return util.ErrorResponse(c, fiber.StatusUnauthorized, unauthorizedMsg+"Student")
	}

	req := &GetStudentByUserIdRequest{
		UsersId: id,
	}

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.StudentService.GetStudentByUserId(c.Context(), req)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Student"+notFoundMsg)
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusOK, "Student successfully retrieved.", res)
}
