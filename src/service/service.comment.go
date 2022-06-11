package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
)

func CreateComment(comment dto.ContentCommentCreator) {
	database.Database.Db.Table("comments").Create(&comment)
}
