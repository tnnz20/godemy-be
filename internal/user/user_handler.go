package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Handler struct {
	Service
	Validate *validator.Validate
}

func NewHandler(s Service, validate *validator.Validate) *Handler {
	return &Handler{
		Service:  s,
		Validate: validate,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserReq
	var isTeacher bool

	roleQuery := c.Query("role")

	if roleQuery == "teacher" {
		isTeacher = true
	} else if roleQuery == "" {
		isTeacher = false
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

	res, err := h.Service.CreateUser(c.Context(), &req, isTeacher)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully.",
		"data":    res,
	})
}

func (h *Handler) GetUserProfileById(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	parseId, _ := uuid.Parse(id)
	req := &GetUserProfileByIdReq{
		ID: parseId,
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID",
		})
	}

	res, err := h.Service.GetUserProfileById(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User ID not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req SignInReq
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

	res, err := h.Service.SignIn(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Success Login",
		"data":    res,
	})
}
