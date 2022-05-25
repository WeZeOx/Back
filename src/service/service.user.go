package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
)

func CreateUser(user dto.User) {
	database.Database.Db.Create(&user)
}

func FindUsers(users []dto.User) []dto.User {
	database.Database.Db.Find(&users)
	return users
}

func FindPost(user dto.User, post []dto.Post) []dto.Post {
	database.Database.Db.Where("user_id = ?", user.ID).Find(&post)
	return post
}

func GetUserByEmail(login dto.Login) dto.User {
	var user dto.User
	database.Database.Db.Where("email = ?", login.Email).Find(&user)
	return user
}

func FindPosts(posts []dto.Post) []dto.Post {
	database.Database.Db.Find(&posts)
	return posts
}

func GetUserById(id string, user dto.User) dto.User {
	database.Database.Db.Find(&user, "id = ?", id)
	return user
}

func GetPostById(id string, post []dto.Post) []dto.Post {
	database.Database.Db.Find(&post, "user_id = ?", id)
	return post
}
