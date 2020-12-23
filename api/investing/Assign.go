package investing

import (
	"invest/redis"
	"invest/structs"

	"github.com/gofiber/fiber/v2"
)

// response herencia de la respuesta general
type response struct {
	structs.Response
	B3 int32 `json:"credit_type_300"`
	B5 int32 `json:"credit_type_500"`
	B7 int32 `json:"credit_type_700"`
}

// Assign propone una opción de asignación de recursos
func Assign(c *fiber.Ctx) error {

	body := c.Context().UserValue("invest").(*structs.InvestBody)
	var asigner structs.Assigner
	var res response
	b3, b5, b7, err := asigner.Assign(body.Investment)
	if err != nil {

		go redis.Store(false, body.String())
		res.BadRequest(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	go redis.Store(true, body.String())
	res.Ok("Ok")
	res.B3 = b3
	res.B5 = b5
	res.B7 = b7
	return c.JSON(res)
}
