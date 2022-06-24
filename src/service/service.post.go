package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/models"
)

func CreateDbPost(post dto.Post) {
	database.Database.Db.Create(&post)
}

func FindPosts() []dto.ResponsePostUser {
	var res []dto.ResponsePostUser
	database.Database.Db.
		Select("*").
		Table("users u").
		Joins("join posts p on u.id = p.user_id").
		Order("p.created_at ASC").
		Scan(&res)
	return res
}

func FindPost(postId string) dto.ResponsePostUser {
	var res dto.ResponsePostUser
	database.Database.Db.
		Table("users").
		Select("*").
		Joins("join posts p on users.id = p.user_id").
		Where("p.post_id = ?", postId).
		Scan(&res)
	return res
}

func GetPostByUserId(id string, post []dto.Post) []dto.Post {
	database.Database.Db.Find(&post, "user_id = ?", id)
	return post
}

func GetPostByPostId(id string, post dto.Post) dto.Post {
	database.Database.Db.Find(&post, "post_id = ?", id)
	return post
}

func DeletePost(post dto.Post) {
	database.Database.Db.
		Where("post_id = ?", post.PostID).
		Delete(&post)
}

func UpdateColumnLike(post dto.Post) {
	database.Database.Db.
		Where("post_id = ?", post.PostID).
		Save(&post)
}

func GetCountCommentsByPost() []int {
	var countResult []int
	database.Database.Db.
		Table("posts p").
		Joins("LEFT JOIN comments c on p.post_id = c.post_id").
		Group("p.post_id").Select("count(c.post_id)").
		Order("p.created_at ASC").
		Scan(&countResult)
	return countResult
}

func GetCountCommentByPost(postId string) int {
	var countResult int
	database.Database.Db.
		Table("posts p").
		Joins("LEFT JOIN comments c on p.post_id = c.post_id").
		Group("p.post_id").
		Where("c.post_id = ?", postId).
		Select("count(c.post_id)").
		Scan(&countResult)
	return countResult
}

func GetPostWithComments(postId string) []dto.PostWithCommentResponse {
	var responseDb []dto.PostWithCommentResponse
	database.Database.Db.
		Table("comments c").
		Select("c.user_id, c.content_comment, c.created_at, u.username, c.like, c.comment_id").
		Joins("join posts p on c.post_id = p.post_id join users u on u.id = c.user_id").
		Order("c.created_at DESC").
		Where("c.post_id = ?", postId).
		Scan(&responseDb)
	return responseDb
}

func DeleteCommentWithPostId(postId string, comment models.Comment) {
	database.Database.Db.
		Table("comments").
		Where("comments.post_id", postId).
		Delete(&comment)
}
