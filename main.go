package main

import (
	"invest/api"
	"invest/config"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/go-redis-connector"
)

func main() {

	app := fiber.New(fiber.Config{
		CaseSensitive:    true,
		ServerHeader:     "Edu",
		DisableKeepalive: true,
	})

	// Configuraciones
	config.Accepts(app)
	config.Security(app)

	appGroup := app.Group("/api")
	api.Router(appGroup)
	// Manejador de Páginas no encontradas
	config.Page404(app)

	// DB Connect
	redis.Connect()

	port := os.Getenv("PORT")
	if port == "" {

		port = "3000"
	}
	app.Listen(":" + port)

	// var creditAsigner structs.Assigner
	// b3, b5, b7, err := creditAsigner.Assign(6701)

	// if err != nil {

	// 	log.Println(err.Error())
	// }

	// log.Println(b3, b5, b7)

}
