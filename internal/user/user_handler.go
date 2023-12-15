package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(err)
	}

	var isTeacher bool
	if c.Query("role") == "teacher" {
		isTeacher = true
	} else {
		isTeacher = false
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := h.Service.CreateUser(c.Context(), &req, isTeacher)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	fmt.Println(res, err)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User has been successfully created",
		"data":    res,
	})

}
