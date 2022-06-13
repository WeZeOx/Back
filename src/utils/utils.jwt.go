package utils

import (
	"Forum-Back-End/src/dto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func CreateJwtToken(user dto.User, isAdmin bool) string {
	godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")
	mySigningKey := []byte(jwtSecret)

	claims := dto.JwtClaims{
		ID:       user.ID,
		Username: user.Username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tk.SignedString(mySigningKey)
	return token
}
