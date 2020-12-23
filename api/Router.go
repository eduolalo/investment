package api

import (
	"invest/api/investing"

	"github.com/gofiber/fiber/v2"
)

// para pasar información ente middlewares, usar c.Fasthttp.SetUserValue y
// c.Fasthttp.UserValue

// Router Es el que agrega las rutas a los paths del api v1
func Router(app fiber.Router) {

	app.Post("/credit-assignment", investing.ValidateBody, investing.Assign)
}
