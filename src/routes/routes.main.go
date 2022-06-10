package routes

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/service"
	"Forum-Back-End/src/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", Welcome)

	UsersRouters(app.Group("/api/users"))
	PostsRouters(app.Group("/api/posts"))
}

func Routes() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/test", func(c *fiber.Ctx) error {
		fmt.Println("la")
		service.GetCountCommentByPost()
		return c.SendString("ma")
	})

	setupRoutes(app)
	utils.AccountAdminExist()

	log.Fatal(app.Listen(":3333"))
}
