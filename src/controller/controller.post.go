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
