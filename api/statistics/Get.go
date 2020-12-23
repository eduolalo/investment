package statistics

import (
	"invest/structs"

	"github.com/gofiber/fiber/v2"
)

type response struct {
	structs.Response
	structs.Statistic
}

// Get obtiene las estadísticas de los requests
func Get(c *fiber.Ctx) error {

	var res response

	res.Calculate()
	res.Ok("")
	return c.JSON(res)
}
