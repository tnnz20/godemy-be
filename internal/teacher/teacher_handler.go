package teacher

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func (h *Handler) GetTeacherIdByUserId(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	role := claims["role"].(string)

	if role != "teacher" {
		return util.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized not a teacher.")
	}
	parseId, _ := uuid.Parse(id)
	req := &GetTeacherIdByUserIdRequest{
		ID: parseId,
	}

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.TeacherService.GetTeacherIdByUserId(c.Context(), req)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Teacher not found.")
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusOK, "Teacher Id successfully retrieved.", res)
}

func (h *Handler) CreateClass(c *fiber.Ctx) error {
	var req CreateClassRequest

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	role := claims["role"].(string)

	if role != "teacher" {
		return util.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized not a teacher.")
	}

	parseId, _ := uuid.Parse(id)
	teacherId := &GetTeacherIdByUserIdRequest{
		ID: parseId,
	}

	// validate UsersId
	if err := h.Validate.Struct(teacherId); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	teacher, err := h.TeacherService.GetTeacherIdByUserId(c.Context(), teacherId)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Teacher not found.")
		}
	}

	if err := c.BodyParser(&req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	req.TeacherId = teacher.ID

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.TeacherService.CreateClass(c.Context(), &req)
	if err != nil {
		util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusCreated, "Class successfully created.", res)
}
