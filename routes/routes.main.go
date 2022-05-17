package routes

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api", Welcome)
}

func Routes() {
	//database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3333"))

}
