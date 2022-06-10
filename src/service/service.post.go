package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"time"
)

func CreateDbPost(post dto.Post) {
	database.Database.Db.Create(&post)
}

func FindPosts() []dto.ResponsePostUser {
	var res []dto.ResponsePostUser
	database.Database.Db.Table("users").Select("*").Joins("join posts p on users.id = p.user_id").Scan(&res)
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
	for i, j := 0, len(countResult)-1; i < j; i, j = i+1, j-1 {
		countResult[i], countResult[j] = countResult[j], countResult[i]
	}
	return countResult
}

type test struct {
	PostId           string    `json:"post_id"`
	UserId           string    `json:"user_id"`
	CommentId        string    `json:"comment_id"`
	CreatedAtComment time.Time `json:"created_at_comment"`
	ContentComment   string    `json:"content_comment"`
	UserID           string    `json:"id"`
	Content          string    `json:"content"`
	CreatedAtPost    time.Time `json:"created_at_post"`
	Like             string    `json:"like"`
	Dislike          string    `json:"dislike"`
	PostID           string    `json:"post_id_post"`
	Category         string    `json:"category"`
	ID               string    `json:"user_id_user"`
	CreatedAtUser    time.Time `json:"created_at_user"`
	Username         string    `json:"username" gorm:"unique"`
	Password         string    `json:"password"`
	Email            string    `json:"email" gorm:"unique"`
}

func GetPostWithComments(postId string) []test {
	var test2 []test

	database.Database.Db.Table("comments c").Select("*").Joins("join posts p on c.post_id = p.post_id join users u on u.id = c.user_id").Where("c.post_id = ?", postId).Scan(&test2)
	return test2
}
