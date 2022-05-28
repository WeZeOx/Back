package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
)

func CreatePost(post dto.Post) {
	database.Database.Db.Create(&post)
}

func FindPosts(posts []dto.Post) []dto.Post {
	database.Database.Db.Find(&posts)
	return posts
}

func GetPostById(id string, post []dto.Post) []dto.Post {
	database.Database.Db.Find(&post, "user_id = ?", id)
	return post
}
