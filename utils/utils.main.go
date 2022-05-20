package utils

import (
	"Forum-Back-End/database"
	"Forum-Back-End/models"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CheckFieldUser(user models.User, array []string) bool {
	var structArray map[string]interface{}
	data, _ := json.Marshal(user)
	json.Unmarshal(data, &structArray)
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

func EmailExist(user models.User) bool {
	var countEmail int64
	database.Database.Db.Where("email = ?", user.Email).Find(&user).Count(&countEmail)
	fmt.Println(countEmail)
	return countEmail > 0
}

func UsernameExist(user models.User) bool {
	var countUsername int64
	database.Database.Db.Where("username = ?", user.Username).Find(&user).Count(&countUsername)
	return countUsername > 0
}
