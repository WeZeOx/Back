package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	userData := c.Locals("user").(dto.User)
	service.CreateUser(userData)
	token := utils.CreateToken(userData)

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		User: userData,
		State: dto.State{
			Message: "Authorized",
			Auth:    true,
			Token:   token,
		},
	})
}

func LoginUser(c *fiber.Ctx) error {
	userData := c.Locals("user").(dto.Login)
	userToLogin := service.GetUserByEmail(userData)

	if utils.CheckPasswordHash(userData.Password, userToLogin.Password) {
		token := utils.CreateToken(userToLogin)

		return c.JSON(dto.Response{
			User: dto.User{
				ID:        userToLogin.ID,
				CreatedAt: userToLogin.CreatedAt,
				Password:  userToLogin.Password,
				Username:  userToLogin.Username,
				Email:     userToLogin.Email,
				Post:      userToLogin.Post,
			},
			State: dto.State{
				Message: "Authenticated",
				Auth:    true,
				Token:   token,
			}})
	}

	return c.Status(fiber.StatusBadRequest).JSON(dto.State{Message: "Email / Password Incorrect", Auth: false})
}

func GetUsers(c *fiber.Ctx) error {
	var posts []dto.Post
	var users []dto.User
	var responseUsersAndPost []dto.User
	users = service.FindUsers(users)

	for _, user := range users {
		posts = service.FindPost(user, posts)
		responseUsersAndPost = append(responseUsersAndPost, utils.CreateResponseUserWithPost(user, posts))
	}
	return c.Status(fiber.StatusOK).JSON(responseUsersAndPost)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id", "")
	var user dto.User
	var post []dto.Post

	user = service.GetUserById(id, user)
	post = service.GetPostById(user.ID, post)

	res := utils.CreateResponseUserWithPost(user, post)

	return c.Status(fiber.StatusOK).JSON(res)
}
