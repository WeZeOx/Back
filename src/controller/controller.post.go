package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetPosts(c *fiber.Ctx) error {
	var responsePostUser []fiber.Map
	var posts []dto.Post
	posts = service.FindPosts(posts)

	for _, post := range posts {
		var user dto.User
		user = service.GetUserById(post.UserID, user)

		responsePostUser = append(responsePostUser, fiber.Map{
			"user": dto.ResponseUser{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Username:  user.Username,
				Email:     user.Email,
			},
			"post": post,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responsePostUser)
}

func CreatePost(c *fiber.Ctx) error {
	postData := c.Locals("post").(dto.Post)
	postData.CreatedAt = time.Now()

	service.CreatePost(postData)

	return c.JSON(fiber.Map{"post": postData})
}
