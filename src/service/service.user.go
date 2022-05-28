package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/models"
)

func CreateUser(user models.User) {
	database.Database.Db.Create(&user)
}

func FindUsers(users []dto.User) []dto.User {
	database.Database.Db.Find(&users)
	return users
}

func GetUserByEmail(login dto.Login) dto.User {
	var user dto.User
	database.Database.Db.Where("email = ?", login.Email).Find(&user)
	return user
}

func GetUserById(id string, user dto.User) dto.User {
	database.Database.Db.Find(&user, "id = ?", id)
	return user
}
