package middleware

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CheckFieldCreateUser(c *fiber.Ctx) error {
	var checkFieldUserArray = []string{"username", "password", "verify_password", "email"}
	var user dto.User
	err := c.BodyParser(&user)
	if (err != nil) ||
		(!utils.CheckFieldUser(user, checkFieldUserArray)) ||
		(user.Password != user.VerifyPassword) ||
		(utils.EmailExist(user)) ||
		(utils.UsernameExist(user)) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.State{Message: "Please check your connexions fields", Auth: false})
	}
	user.ID = uuid.New().String()
	user.Password, _ = utils.HashPassword(user.Password)
	c.Locals("user", user)

	return c.Next()
}

func CheckFieldLogin(c *fiber.Ctx) error {
	fmt.Println("here")
	var checkFieldLoginArray = []string{"email", "password"}
	var login dto.Login
	err := c.BodyParser(&login)
	if (!utils.CheckFieldLogin(login, checkFieldLoginArray)) ||
		err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.State{Message: "Missing fields", Auth: false})
	}
	c.Locals("user", login)
	return c.Next()
}
