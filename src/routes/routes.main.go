package routes

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", Welcome)

	app.Get("/test", func(c *fiber.Ctx) error {
		var res []dto.ResponsePostUser
		database.Database.Db.Table("users").Select("*").Joins("join posts p on users.id = p.user_id").Scan(&res)
		return c.JSON(res)
	})

	UsersRouters(app.Group("/api/users"))
	PostsRouters(app.Group("/api/posts"))
}

func Routes() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)
	log.Fatal(app.Listen(":3333"))
}
