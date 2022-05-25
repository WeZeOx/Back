package service

import (
	"Forum-Back-End/src/database"
	"Forum-Back-End/src/dto"
)

func CreatePost(post dto.Post) {
	database.Database.Db.Create(&post)
}
