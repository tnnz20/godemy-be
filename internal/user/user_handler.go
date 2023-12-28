package user

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/util"
)

type Handler struct {
	UserService Service
	Validate    *validator.Validate
}

func NewHandler(s Service, validate *validator.Validate) *Handler {
	return &Handler{
		UserService: s,
		Validate:    validate,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	roleQuery := c.Query("role")

	if roleQuery == "teacher" {
		req.Role = "teacher"
	} else if roleQuery == "" {
		req.Role = "student"
	} else {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Role.")
	}

	if err := c.BodyParser(&req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if user, err := h.UserService.GetUserByEmail(c.Context(), &req.Email); err != sql.ErrNoRows && user.Email == req.Email {
		return util.ErrorResponse(c, fiber.StatusConflict, "Email already exists.")
	}

	res, err := h.UserService.CreateUser(c.Context(), &req)
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusCreated, "User successfully created.", res)
}

func (h *Handler) GetUserProfileById(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	parseId, _ := uuid.Parse(id)
	req := &GetUserProfileByIdRequest{
		ID: parseId,
	}

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Id.")
	}

	res, err := h.UserService.GetUserProfileById(c.Context(), req)
	if err != nil {
		if err == sql.ErrNoRows {
			util.ErrorResponse(c, fiber.StatusNotFound, "User Id not found.")
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusOK, "User profile successfully retrieved.", res)
}
