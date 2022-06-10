package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
)

func CreateDbPost(post dto.Post) {
	database.Database.Db.Create(&post)
}

func FindPosts() []dto.ResponsePostUser {
	var res []dto.ResponsePostUser
	database.Database.Db.Table("users").Select("*").Joins("join posts p on users.id = p.user_id").Scan(&res)
	return res
}

func GetPostById(id string, post []dto.Post) []dto.Post {
	database.Database.Db.Find(&post, "user_id = ?", id)
	return post
}

func GetPostByPostId(id string, post dto.Post) dto.Post {
	database.Database.Db.Find(&post, "post_id = ?", id)
	return post
}

func DeletePost(post dto.Post) {
	database.Database.Db.Where("post_id = ?", post.PostID).Delete(&post)
}

func UpdateColumnLike(post dto.Post) {
	database.Database.Db.Where("post_id = ?", post.PostID).Save(&post)
}
