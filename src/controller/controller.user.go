package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	userData := c.Locals("user").(dto.User)
	user := utils.CreateDbUser(userData)

	service.CreateUser(user)
	token := utils.CreateToken(userData)

	return c.Status(200).JSON(fiber.Map{
		"user": dto.ResponseUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Username:  user.Username,
			Email:     user.Email,
		},
		"state": dto.State{
			Message: "Authorized",
			Auth:    true,
			Token:   token,
		}})
}

func LoginUser(c *fiber.Ctx) error {
	userData := c.Locals("user").(dto.Login)
	userToLogin := service.GetUserByEmail(userData)
	user := utils.CreateDbUser(userToLogin)

	if utils.CheckPasswordHash(userData.Password, userToLogin.Password) {
		token := utils.CreateToken(userToLogin)

		return c.Status(200).JSON(fiber.Map{
			"user": dto.ResponseUser{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Username:  user.Username,
				Email:     user.Email,
			},
			"state": dto.State{
				Message: "Authorized",
				Auth:    true,
				Token:   token,
			}})
	}

	return c.Status(fiber.StatusBadRequest).JSON(dto.State{Message: "Email or Password Incorrect", Auth: false})
}

func GetUsers(c *fiber.Ctx) error {
	var posts []dto.Post
	var users []dto.User
	var responseUsersAndPost []fiber.Map
	users = service.FindUsers(users)

	for _, user := range users {
		posts = service.GetPostById(user.ID, posts)
		responseUsersAndPost = append(responseUsersAndPost, fiber.Map{
			"user": dto.ResponseUser{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Username:  user.Username,
				Email:     user.Email,
			},
			"post": posts,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responseUsersAndPost)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id", "")
	var user dto.User
	var post []dto.Post

	user = service.GetUserById(id, user)
	post = service.GetPostById(user.ID, post)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"post": post,
		"user": dto.ResponseUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Username:  user.Username,
			Email:     user.Email,
		},
	})
}
