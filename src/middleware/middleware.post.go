package middleware

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
)

func CheckToken(c *fiber.Ctx) error {
	tokenString := c.GetReqHeaders()["Authorization"]
	godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")

	type Claims struct {
		ID string `json:"id"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.State{
			Message: "Wrong token",
			Auth:    false,
		})
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.State{
			Message: "Wrong token",
			Auth:    false,
		})
	}
}

func CheckFieldCreatePost(c *fiber.Ctx) error {
	var checkFieldPostArray = []string{"id", "content"}
	var post dto.Post
	err := c.BodyParser(&post)

	if (err != nil) ||
		!utils.CheckFieldPost(post, checkFieldPostArray) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.State{
			Message: "Bad Fields",
			Auth:    false,
		})
	}

	post.PostID = uuid.New().String()
	c.Locals("post", post)

	return c.Next()
}
