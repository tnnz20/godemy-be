package user

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

// TODO: Update Response Handler
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	roleQuery := c.Query("role")

	if roleQuery == "teacher" {
		req.Role = "teacher"
	} else if roleQuery == "" {
		req.Role = "student"
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Role",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	if err := h.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if user, err := h.UserService.GetUserByEmail(c.Context(), &req.Email); err != sql.ErrNoRows && user.Email == req.Email {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Email already exists",
		})
	}

	res, err := h.UserService.CreateUser(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully registered.",
		"data":    res,
	})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID",
		})
	}

	res, err := h.UserService.GetUserProfileById(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User ID not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User profile successfully retrieved.",
		"data":    res,
	})
}
