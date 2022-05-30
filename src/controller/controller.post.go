package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetPosts(c *fiber.Ctx) error {
	posts := service.FindPosts()
	return c.Status(fiber.StatusOK).JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	postData := c.Locals("post").(dto.Post)
	postData.CreatedAt = time.Now()
	service.CreatePost(postData)
	return c.JSON(fiber.Map{"post": postData})
}
