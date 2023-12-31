package teacher

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tnnz20/godemy-be/util"
)

type Handler struct {
	TeacherService Service
	Validate       *validator.Validate
}

func NewHandler(s Service, validate *validator.Validate) *Handler {
	return &Handler{
		TeacherService: s,
		Validate:       validate,
	}
}

const (
	unauthorizedMsg = "Unauthorized not a "
	notFoundMsg     = " not found."
)

func (h *Handler) GetTeacherIdByUserId(c *fiber.Ctx) error {
	id, role := util.JwtClaim(c)

	if role != "teacher" {
		return util.ErrorResponse(c, fiber.StatusUnauthorized, unauthorizedMsg+"Teacher")
	}

	req := &GetTeacherIdByUserIdRequest{
		UserId: id,
	}

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.TeacherService.GetTeacherIdByUserId(c.Context(), req)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Teacher"+notFoundMsg)
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusOK, "Teacher Id successfully retrieved.", res)
}

func (h *Handler) GetAllClassByTeacherId(c *fiber.Ctx) error {
	id, role := util.JwtClaim(c)

	if role != "teacher" {
		return util.ErrorResponse(c, fiber.StatusUnauthorized, unauthorizedMsg+"Teacher")
	}

	UserId := &GetTeacherIdByUserIdRequest{
		UserId: id,
	}

	if err := h.Validate.Struct(UserId); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	teacher, err := h.TeacherService.GetTeacherIdByUserId(c.Context(), UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Teacher"+notFoundMsg)
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	req := &GetClassByTeacherIdRequest{
		TeacherId: teacher.ID,
	}

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.TeacherService.GetAllClassByTeacherId(c.Context(), req)
	if err != nil {
		if err.Error() == "null" {
			return util.ErrorResponse(c, fiber.StatusNotFound, "No classes found from given teacher id.")
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusOK, "Classes successfully retrieved.", res)
}
