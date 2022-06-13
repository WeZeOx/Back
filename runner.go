package main

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/routes"
	"Forum-Back-End/src/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", Welcome)
	routes.UsersRouters(app.Group("/api/users"))
	routes.PostsRouters(app.Group("/api/posts"))
	routes.CommentsRouters(app.Group("/api/comments"))
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)
	service.AccountAdminExist()

	log.Fatal(app.Listen(":3333"))
}
