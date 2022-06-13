package middleware

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"regexp"
)

func CheckFieldCreateUser(c *fiber.Ctx) error {
	var checkFieldUserArray = []string{"username", "password", "verify_password", "email"}
	var user dto.User
	err := c.BodyParser(&user)
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseState{Message: "Your email must be valid", Auth: false})
	}
	if service.EmailExist(user) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseState{Message: "Email already taken", Auth: false})
	}
	if service.UsernameExist(user) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseState{Message: "Username already taken", Auth: false})
	}
	if (err != nil) ||
		(!utils.CheckFieldUser(user, checkFieldUserArray)) ||
		(user.Password != user.VerifyPassword) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseState{Message: "Please check your connexions fields", Auth: false})
	}
	user.ID = uuid.New().String()
	user.Password, _ = utils.HashPassword(user.Password)
	c.Locals("user", user)
	return c.Next()
}

func CheckFieldLogin(c *fiber.Ctx) error {
	var checkFieldLoginArray = []string{"email", "password"}
	var login dto.BodyLoginRequest
	err := c.BodyParser(&login)
	if (!utils.CheckFieldLogin(login, checkFieldLoginArray)) ||
		err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseState{Message: "Missing fields", Auth: false})
	}
	c.Locals("user", login)
	return c.Next()
}
