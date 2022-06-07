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
	user := utils.CreateDbUser(userData)

	service.CreateUser(user)
	token := utils.CreateToken(userData, false)

	return c.Status(200).JSON(fiber.Map{
		"user": dto.ResponseUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Username:  user.Username,
		},
		"state": dto.State{
			Message: "Authorized",
			Auth:    true,
			Token:   token,
		}})
}

func LoginUser(c *fiber.Ctx) error {
	godotenv.Load(".env")
	userData := c.Locals("user").(dto.Login)
	userToLogin := service.GetUserByEmail(userData)
	user := utils.CreateDbUser(userToLogin)
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	ADMIN_PASSWORD := os.Getenv("ADMIN_PASSWORD")

	if utils.CheckPasswordHash(userData.Password, userToLogin.Password) &&
		userData.Password == ADMIN_PASSWORD &&
		userData.Email == ADMIN_EMAIL {
		token := utils.CreateToken(userToLogin, true)

		return c.Status(200).JSON(fiber.Map{
			"user": dto.ResponseUser{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Username:  user.Username,
			},
			"state": dto.State{
				Message: "Authorized",
				Auth:    true,
				Token:   token,
			}})
	}

	if utils.CheckPasswordHash(userData.Password, userToLogin.Password) {
		token := utils.CreateToken(userToLogin, false)

		return c.Status(200).JSON(fiber.Map{
			"user": dto.ResponseUser{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Username:  user.Username,
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
			},
			"post": posts,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responseUsersAndPost)
}

func GetUser(c *fiber.Ctx) error {
	godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")

	id := c.Params("id", "")
	var user dto.User
	var post []dto.Post

	userAdmin := service.GetAdminUserByEmail(ADMIN_EMAIL)
	user = service.GetUserById(id, user)
	post = service.GetPostById(user.ID, post)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"post": post,
		"user": dto.ResponseUser{
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

	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return c.JSON(fiber.Map{"isAdmin": false})
	}

	cls, _ := token.Claims.(*dto.Claims)
	return c.JSON(fiber.Map{"isAdmin": cls.IsAdmin})
}
