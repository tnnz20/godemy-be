package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SuccessCreated(c *fiber.Ctx) error {
	return c.Status(http.StatusCreated).JSON(SuccessCreatedMessage{
		Code:    http.StatusCreated,
		Message: StatusSuccessCreatedName,
	})
}

func SuccessOK(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(SuccessOKMessage{
		Code:    http.StatusOK,
		Message: StatusSuccessOK,
		Data:    data,
	})
}

func ErrorBadRequest(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusBadRequest).JSON(ErrorMessage{
		Code: http.StatusBadRequest,
		Error: &ErrorFormat{
			ErrorName:        StatusBadRequestErrorName,
			ErrorDescription: err.Error(),
		},
	})
}

func ErrorNotFound(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusNotFound).JSON(ErrorMessage{
		Code: http.StatusNotFound,
		Error: &ErrorFormat{
			ErrorName:        StatusNotFound,
			ErrorDescription: err.Error(),
		},
	})
}

func ErrorConflict(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusConflict).JSON(ErrorMessage{
		Code: http.StatusConflict,
		Error: &ErrorFormat{
			ErrorName:        StatusConflict,
			ErrorDescription: err.Error(),
		},
	})
}

func ErrorUnauthorized(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(ErrorMessage{
		Code: http.StatusUnauthorized,
		Error: &ErrorFormat{
			ErrorName:        StatusUnauthorizedErrorName,
			ErrorDescription: StatusUnauthorizedErrorDescription,
		},
	})
}

func ErrorInvalidToken(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(ErrorMessage{
		Code: http.StatusUnauthorized,
		Error: &ErrorFormat{
			ErrorName:        StatusUnauthorizedErrorName,
			ErrorDescription: StatusUnauthorizedInvalidToken,
		},
	})
}

func ErrorForbiddenAccess(c *fiber.Ctx) error {
	return c.Status(http.StatusForbidden).JSON(ErrorMessage{
		Code: http.StatusForbidden,
		Error: &ErrorFormat{
			ErrorName:        StatusForbidden,
			ErrorDescription: StatusForbiddenDescription,
		},
	})

}

func ErrorValidateToken(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusUnauthorized).JSON(ErrorMessage{
		Code: http.StatusUnauthorized,
		Error: &ErrorFormat{
			ErrorName:        StatusUnauthorizedErrorName,
			ErrorDescription: err.Error(),
		},
	})

}

func InternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusInternalServerError).JSON(ErrorMessage{
		Code: http.StatusInternalServerError,
		Error: &ErrorFormat{
			ErrorName:        StatusInternalServerName,
			ErrorDescription: err.Error(),
		},
	})
}
