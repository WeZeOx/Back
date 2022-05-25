package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []dto.Post
	posts = service.FindPosts(posts)
	return c.Status(fiber.StatusOK).JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	postData := c.Locals("post").(dto.Post)
	service.CreatePost(postData)

	return c.JSON(postData)
}
