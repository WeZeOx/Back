package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/models"
	"Forum-Back-End/src/utils"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func CreateUserInDb(user models.User) {
	database.Database.Db.Create(&user)
}

func GetUserByEmail(email string) dto.User {
	var user dto.User
	database.Database.Db.Table("users").Where("email = ?", email).Scan(&user)
	return user
}

func GetUserById(id string, user dto.User) dto.User {
	database.Database.Db.Find(&user, "id = ?", id)
	return user
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

func AccountAdminExist() bool {
	var user dto.User
	ADMIN_EMAIL := utils.OpenDotEnvAndQueryTheValue("ADMIN_EMAIL")
	ADMIN_PASSWORD := utils.OpenDotEnvAndQueryTheValue("ADMIN_PASSWORD")
	ADMIN_PASSWORD_HASH, _ := utils.HashPassword(ADMIN_PASSWORD)
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

		userDB := utils.CreateDbUserSchema(user)
		CreateUserInDb(userDB)
	}
	return true
}
