package utils

import (
	"Forum-Back-End/src/dto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func CreateToken(user dto.User, isAdmin bool) string {
	godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")
	mySigningKey := []byte(jwtSecret)

	claims := dto.Claims{
		ID:      user.ID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString(mySigningKey)
	return token
}
