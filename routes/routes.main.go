package routes

import (
	"Forum-Back-End/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api", Welcome)
}

func Routes() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)
	log.Fatal(app.Listen(":3333"))
}
