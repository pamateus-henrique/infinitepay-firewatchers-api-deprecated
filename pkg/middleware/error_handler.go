package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := ""

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	if message == "" && code > 399 && code < 500 {
		message = "Something is wrong on your request"
	}

	if message == "" && code > 499 {
		message = "Something went wrong while processing your request, this is not your fault"
	}

	// Send custom error page
	err = c.Status(code).JSON(fiber.Map{
		"error": true,
		"message": message,
	})
	if err != nil {
		// In case the SendFile fails
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}