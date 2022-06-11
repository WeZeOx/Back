package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"sort"
)

func CreateDbPost(post dto.Post) {
	database.Database.Db.Create(&post)
}

func FindPosts() []dto.ResponsePostUser {
	var res []dto.ResponsePostUser
	database.Database.Db.Table("users").Select("*").Joins("join posts p on users.id = p.user_id").Scan(&res)
	return res
}

func FindPost(postId string) dto.ResponsePostUser {
	var res dto.ResponsePostUser
	database.Database.Db.Table("users").Select("*").Joins("join posts p on users.id = p.user_id").Where("p.post_id = ?", postId).Scan(&res)
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
	database.Database.Db.Where("post_id = ?", post.PostID).Delete(&post)
}

func UpdateColumnLike(post dto.Post) {
	database.Database.Db.Where("post_id = ?", post.PostID).Save(&post)
}

func GetCountCommentByPost() []int {
	var countResult []int
	database.Database.Db.Table("posts p").Joins("LEFT JOIN comments c on p.post_id = c.post_id").Group("p.post_id").Select("count(c.post_id)").Order("p.created_at DESC").Scan(&countResult)
	sort.Slice(countResult, func(i, j int) bool {
		return true
	})
	return countResult
}

func GetPostWithComments(postId string) []dto.PostWithCommentResponse {
	var responseDb []dto.PostWithCommentResponse
	database.Database.Db.Table("comments c").Select("c.user_id, c.content_comment,p.content, p.created_at, p.like, p.category, u.username").Joins("join posts p on c.post_id = p.post_id join users u on u.id = c.user_id").Where("c.post_id = ?", postId).Scan(&responseDb)
	return responseDb
}
