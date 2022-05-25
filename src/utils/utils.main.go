package utils

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func CheckFieldUser(user dto.User, array []string) bool {
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

func CheckFieldLogin(user dto.Login, array []string) bool {
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

func CheckFieldPost(user dto.Post, array []string) bool {
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

func EmailExist(user dto.User) bool {
	var countEmail int64
	database.Database.Db.Where("email = ?", user.Email).Find(&user).Count(&countEmail)
	return countEmail > 0
}

func UsernameExist(user dto.User) bool {
	var countUsername int64
	database.Database.Db.Where("username = ?", user.Username).Find(&user).Count(&countUsername)
	return countUsername > 0
}

func CreateToken(user dto.User) string {
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

func CreateResponseUserWithPost(user dto.User, PostArr []dto.Post) dto.User {
	return dto.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		Password:  user.Password,
		Email:     user.Email,
		Post:      PostArr,
	}
}
