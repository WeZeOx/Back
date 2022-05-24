package utils

import (
	"Forum-Back-End/database"
	"Forum-Back-End/structures"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func CheckFieldUser(user structures.User, array []string) bool {
	var structArray map[string]interface{}
	data, _ := json.Marshal(user)
	err := json.Unmarshal(data, &structArray)
	if err != nil {
		return false
	}

	for _, item := range array {
		if structArray[item] == "" || structArray[item] == nil {
			return false
		}
	}
	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EmailExist(user structures.User) bool {
	var countEmail int64
	database.Database.Db.Where("email = ?", user.Email).Find(&user).Count(&countEmail)
	return countEmail > 0
}

func UsernameExist(user structures.User) bool {
	var countUsername int64
	database.Database.Db.Where("username = ?", user.Username).Find(&user).Count(&countUsername)
	return countUsername > 0
}

func CreateToken(user structures.User) string {
	godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")
	mySigningKey := []byte(jwtSecret)

	type Claims struct {
		ID string `json:"id"`
		jwt.RegisteredClaims
	}
	claims := Claims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString(mySigningKey)
	return token
}

func CheckToken(c *fiber.Ctx) error {
	tokenString := c.GetReqHeaders()["Authorization"]
	godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")

	type Claims struct {
		ID string `json:"id"`
		jwt.RegisteredClaims
	}

	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(structures.State{
			Message: "Wrong token",
			Auth:    false,
		})
	}
}
