package utils

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateDbUserSchema(userData dto.User) models.User {
	return models.User{
		ID:        userData.ID,
		CreatedAt: userData.CreatedAt,
		Username:  userData.Username,
		Password:  userData.Password,
		Email:     userData.Email,
	}
}

func CreateUserPostResponse(postData dto.ResponsePostUser, isAdmin bool, num int) dto.PostModel {
	return dto.PostModel{
		UserID:          postData.Post.UserID,
		CreatedAt:       postData.User.CreatedAt,
		Username:        postData.Username,
		Content:         postData.Content,
		Like:            postData.Like,
		PostID:          postData.PostID,
		Categories:      postData.Post.Category,
		Admin:           isAdmin,
		NumberOfComment: num,
	}
}

func CreatePostResponse(post dto.Post, user dto.User, isAdmin bool, numberOfComment int) dto.PostModel {
	return dto.PostModel{
		UserID:          user.ID,
		CreatedAt:       post.CreatedAt,
		Username:        user.Username,
		Content:         post.Content,
		Like:            post.Like,
		PostID:          post.PostID,
		Categories:      post.Category,
		Admin:           isAdmin,
		NumberOfComment: numberOfComment,
	}
}

func CreateSuccessfulLoginResponse(user models.User, token, message string, auth bool) fiber.Map {
	return fiber.Map{
		"user": dto.ResponseWithSafeField{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Username:  user.Username,
		},
		"state": dto.ResponseState{
			Message: message,
			Auth:    auth,
			Token:   token,
		}}
}
