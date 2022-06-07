package utils

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/models"
)

func CreateDbUser(userData dto.User) models.User {
	return models.User{
		ID:        userData.ID,
		CreatedAt: userData.CreatedAt,
		Username:  userData.Username,
		Password:  userData.Password,
		Email:     userData.Email,
	}
}

func CreateUserPostResponse(postData dto.ResponsePostUser, isAdmin bool) dto.PostUserResponseForFront {
	return dto.PostUserResponseForFront{
		ID:        postData.Post.UserID,
		CreatedAt: postData.User.CreatedAt,
		Username:  postData.Username,
		Content:   postData.Content,
		Like:      postData.Like,
		PostID:    postData.PostID,
		Category:  postData.Post.Category,
		Admin:     isAdmin,
	}
}
