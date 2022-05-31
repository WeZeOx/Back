package utils

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
	"time"
)

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

func AccountAdminExist() bool {
	godotenv.Load(".env")
	var user dto.User

	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	ADMIN_PASSWORD := os.Getenv("ADMIN_PASSWORD")
	ADMIN_PASSWORD_HASH, _ := HashPassword(ADMIN_PASSWORD)
	user.Email = ADMIN_EMAIL

	if EmailExist(user) {
		fmt.Println("\nAdmin user already exist")
		return true
	} else {
		fmt.Println("\nAdmin user created")
		user.ID = uuid.New().String()
		user.Password = ADMIN_PASSWORD_HASH
		user.Username = "ADMIN ACCOUNT"
		user.CreatedAt = time.Now()

		userDB := CreateDbUser(user)
		service.CreateUser(userDB)
	}
	return true
}
