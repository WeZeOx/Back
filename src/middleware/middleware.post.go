package middleware

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
)

func DecodeToken(c *fiber.Ctx) error {
	tokenHeader := c.Locals("token").(*jwt.Token)
	_ = godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")

	var AccessToken map[string]string
	stringify, _ := json.Marshal(&tokenHeader)
	_ = json.Unmarshal(stringify, &AccessToken)

	token, _ := jwt.ParseWithClaims(AccessToken["Raw"], &dto.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	decodedToken := token.Claims.(*dto.JwtClaims)
	c.Locals("decodedToken", decodedToken)

	return c.Next()
}

func CheckToken(c *fiber.Ctx) error {
	tokenString := c.GetReqHeaders()["Authorization"]
	_ = godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &dto.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ResponseState{
			Message: "Wrong token",
			Auth:    false,
		})
	}

	if _, ok := token.Claims.(*dto.JwtClaims); ok && token.Valid {
		c.Locals("token", token)
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ResponseState{
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
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseState{
			Message: "Bad Fields",
			Auth:    false,
		})
	}

	post.PostID = uuid.New().String()
	c.Locals("post", post)
	return c.Next()
}
