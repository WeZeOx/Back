package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/models"
)

func CreateComment(comment dto.ContentCommentCreator) {
	database.Database.Db.Table("comments").Create(&comment)
}

func GetCommentByCommentId(commentId string) models.Comment {
	var comment models.Comment
	database.Database.Db.
		Table("comments").
		Where("comment_id = ?", commentId).
		Scan(&comment)
	return comment
}

func SaveLikeColumn(comment models.Comment) {
	database.Database.Db.
		Where("comment_id = ?", comment.CommentId).
		Save(&comment)
}

func DeleteComment(commentId string) {
	var comment models.Comment
	comment.CommentId = commentId
	database.Database.Db.
		Table("comments").
		Where("comment_id = ?", commentId).
		Delete(&comment)
}
