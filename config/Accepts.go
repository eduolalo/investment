package config

import (
	"invest/structs"

	"strings"

	"github.com/gofiber/fiber/v2"
)

// Accepts agraga configuraciones para las de los requests
func Accepts(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {

		// Sólamente se aceptarán los Content-Type application/json
		c.Accepts("application/json")

		// Compresión al máximo
		c.AcceptsEncodings("compress", "br")
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {

		// Sólamente se aceptarán los Content-Type application/json
		ct := c.Context().Request.Header.ContentType()
		// ct := c.Request.Header.ContentType()
		if !strings.Contains(string(ct), "application/json") {

			var res structs.Response
			res.BadRequest("Content-Type erroneo.")
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}

		return c.Next()
	})

}
