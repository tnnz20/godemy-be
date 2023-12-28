package util

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}

func SuccessResponse(c *fiber.Ctx, statusCode int, message, data interface{}) error {
	if data != nil {
		return c.Status(statusCode).JSON(fiber.Map{
			"status":  "success",
			"message": message,
			"data":    data,
		})
	}
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "success",
		"message": message,
	})
}
