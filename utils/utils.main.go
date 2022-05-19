package utils

import (
	"Forum-Back-End/models"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

func CheckField(user models.User, array []string) bool {
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
