package routes

import (
	"Forum-Back-End/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", Welcome)
	app.Get("/api/users", GetUsers)
	app.Get("/api/:id", GetUser)

	app.Post("/api/signup", CreateUser)

	app.Post("/api/signin", LoginUser)
}

func Routes() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)
	log.Fatal(app.Listen(":3333"))
}
