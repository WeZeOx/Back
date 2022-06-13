package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"os"
)

func CreateUser(c *fiber.Ctx) error {
	userData := c.Locals("user").(dto.User)
	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")

	user := utils.CreateDbUserSchema(userData)
	service.CreateUserInDb(user)
	token := utils.CreateJwtToken(userData, userData.Email == ADMIN_EMAIL)

	return c.Status(200).JSON(utils.CreateSuccessfulLoginResponse(user, token, "Authorized", userData.Email == ADMIN_EMAIL))
}

func LoginUser(c *fiber.Ctx) error {
	godotenv.Load(".env")
	userData := c.Locals("user").(dto.BodyLoginRequest)
	userToLogin := service.GetUserByEmail(userData.Email)
	user := utils.CreateDbUserSchema(userToLogin)
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	ADMIN_PASSWORD := os.Getenv("ADMIN_PASSWORD")

	if utils.CheckPasswordHash(userData.Password, userToLogin.Password) &&
		userData.Password == ADMIN_PASSWORD &&
		userData.Email == ADMIN_EMAIL {
		token := utils.CreateJwtToken(userToLogin, true)
		return c.Status(200).JSON(utils.CreateSuccessfulLoginResponse(user, token, "Authorized", true))
	}
	if utils.CheckPasswordHash(userData.Password, userToLogin.Password) {
		token := utils.CreateJwtToken(userToLogin, false)
		return c.Status(200).JSON(utils.CreateSuccessfulLoginResponse(user, token, "Authorized", true))
	}
	return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseState{Message: "Email or Password Incorrect", Auth: false})
}

func GetUser(c *fiber.Ctx) error {
	godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	id := c.Params("userId")
	var user dto.User
	var post []dto.Post

	post = service.GetPostByUserId(user.ID, post)
	userAdmin := service.GetUserByEmail(ADMIN_EMAIL)
	user = service.GetUserById(id, user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"post": post,
		"user": dto.ResponseWithSafeField{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Username:  user.Username,
		},
		"admin": user.ID == userAdmin.ID,
	})
}

func UserIsAdmin(c *fiber.Ctx) error {
	tokenString := c.GetReqHeaders()["Authorization"]
	godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &dto.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return c.JSON(fiber.Map{"isAdmin": false})
	}

	cls, _ := token.Claims.(*dto.JwtClaims)
	return c.JSON(fiber.Map{"isAdmin": cls.IsAdmin})
}
