package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/models"
	"Forum-Back-End/src/service"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"time"
)

func GetPosts(c *fiber.Ctx) error {
	posts := service.FindPosts()
	var res []dto.PostUserResponseForFront
	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	adminSchema := service.GetAdminUserByEmail(ADMIN_EMAIL)
	arrCom := service.GetCountCommentsByPost()

	for idx, post := range posts {
		isAdmin := adminSchema.ID == post.Post.UserID
		res = append(res, utils.CreateUserPostResponse(post, isAdmin, arrCom[idx]))
	}

	//c.Set("content-encoding", "gzip")
	return c.Status(fiber.StatusOK).JSON(res)
}

func CreatePost(c *fiber.Ctx) error {
	postData := c.Locals("post").(dto.Post)
	token := c.Locals("decodedToken").(*dto.Claims)
	postData.CreatedAt = time.Now()
	service.CreateDbPost(postData)
	return c.JSON(fiber.Map{"post": postData, "admin": token.IsAdmin})
}

func DeletePost(c *fiber.Ctx) error {
	var post dto.Post
	var comment models.Comment
	postId := c.Params("postId")

	comment.PostId = postId
	post.PostID = postId

	service.DeletePost(post)

	service.DeleteCommentWithPostId(postId, comment)

	return c.JSON(fiber.Map{
		"isOk": true,
	})
}

func UnlikePost(c *fiber.Ctx) error {
	var post dto.Post
	var user dto.User
	postId := c.Params("postId")
	decodedToken := c.Locals("decodedToken").(*dto.Claims)
	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	adminSchema := service.GetAdminUserByEmail(ADMIN_EMAIL)

	post = service.GetPostByPostId(postId, post)
	userId := decodedToken.ID
	newLikeColumn := ""

	userWhoLikeArr := strings.Split(post.Like, ",")
	user = service.GetUserById(post.UserID, user)

	for _, id := range userWhoLikeArr {
		if id != userId {
			newLikeColumn += id + ","
		}
	}

	newLikeColumn = newLikeColumn[:len(newLikeColumn)-1]
	post.Like = newLikeColumn
	service.UpdateColumnLike(post)
	numberOfComment := service.GetCountCommentByPost(postId)

	return c.JSON(dto.PostUserResponseForFront{
		UserID:       user.ID,
		CreatedAt:    post.CreatedAt,
		Username:     user.Username,
		Content:      post.Content,
		Like:         post.Like,
		PostID:       post.PostID,
		Categories:   post.Category,
		Admin:        adminSchema.ID == user.ID,
		NumberOfPost: numberOfComment,
	})
}

func LikePost(c *fiber.Ctx) error {
	var post dto.Post
	var user dto.User
	postId := c.Params("postId")
	decodedToken := c.Locals("decodedToken").(*dto.Claims)

	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	adminSchema := service.GetAdminUserByEmail(ADMIN_EMAIL)

	post = service.GetPostByPostId(postId, post)
	user = service.GetUserById(post.UserID, user)
	numberOfComment := service.GetCountCommentByPost(postId)
	userId := decodedToken.ID

	post.Like += userId + ","
	service.UpdateColumnLike(post)

	return c.JSON(dto.PostUserResponseForFront{
		UserID:       user.ID,
		CreatedAt:    post.CreatedAt,
		Username:     user.Username,
		Content:      post.Content,
		Like:         post.Like,
		PostID:       post.PostID,
		Categories:   post.Category,
		Admin:        adminSchema.ID == user.ID,
		NumberOfPost: numberOfComment,
	})
}
