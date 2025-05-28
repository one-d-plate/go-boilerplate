package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/go-boilerplate.git/src/pkg"
)

func HandleFiberError(c *fiber.Ctx, err error) error {
	var fiberErr *fiber.Error
	pkg.LogError("Error: ", err)
	if errors.As(err, &fiberErr) {
		// Jika sudah berupa Fiber error, gunakan status dan message-nya
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"message": fiberErr.Message,
			"data":    nil,
		})
	}

	// Default: Internal Server Error
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal server error",
		"data":    nil,
	})
}
