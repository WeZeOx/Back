package routes

import (
	"Forum-Back-End/database"
	"Forum-Back-End/structures"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var post structures.Post

	if err := c.BodyParser(&post); err != nil {
		c.Status(400).JSON(structures.State{Message: "Missing fields", Auth: false})
	}

	database.Database.Db.Create(&post)

	return c.JSON(post)
}
