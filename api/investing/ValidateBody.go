package investing

import (
	"invest/structs"

	"github.com/gofiber/fiber/v2"
)

// ValidateBody revisa que el body sea correcto
func ValidateBody(c *fiber.Ctx) error {

	var body structs.InvestBody

	body.Unmarshal(c.Body())
	if body.Investment == 0 {

		var res structs.Response
		res.BadRequest("La inversión debe ser mayor a $300")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	c.Context().SetUserValue("invest", &body)
	return c.Next()
}
