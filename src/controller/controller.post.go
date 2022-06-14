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
	var res []dto.PostModel
	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	adminSchema := service.GetUserByEmail(ADMIN_EMAIL)
	arrCom := service.GetCountCommentsByPost()

	for idx, post := range posts {
		isAdmin := adminSchema.ID == post.Post.UserID
		res = append(res, utils.CreateUserPostResponse(post, isAdmin, arrCom[idx]))
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func CreatePost(c *fiber.Ctx) error {
	postData := c.Locals("post").(dto.Post)
	decodedToken := c.Locals("decodedToken").(*dto.JwtClaims)
	postData.CreatedAt = time.Now()
	service.CreateDbPost(postData)

	return c.JSON(utils.CreatePostResponse(postData, decodedToken.Username, decodedToken.ID, decodedToken.IsAdmin, 0))
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
	var newLikeColumn string
	postId := c.Params("postId")
	decodedToken := c.Locals("decodedToken").(*dto.JwtClaims)
	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	adminSchema := service.GetUserByEmail(ADMIN_EMAIL)

	post = service.GetPostByPostId(postId, post)

	userWhoLikeArr := strings.Split(post.Like, ",")

	for _, id := range userWhoLikeArr {
		if id != decodedToken.ID {
			newLikeColumn += id + ","
		}
	}

	newLikeColumn = newLikeColumn[:len(newLikeColumn)-1]
	post.Like = newLikeColumn
	service.UpdateColumnLike(post)
	numberOfComment := service.GetCountCommentByPost(postId)
	isAdmin := adminSchema.ID == decodedToken.ID

	return c.JSON(utils.CreatePostResponse(post, decodedToken.Username, decodedToken.ID, isAdmin, numberOfComment))
}

func LikePost(c *fiber.Ctx) error {
	var post dto.Post
	postId := c.Params("postId")
	decodedToken := c.Locals("decodedToken").(*dto.JwtClaims)

	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")

	adminSchema := service.GetUserByEmail(ADMIN_EMAIL)
	post = service.GetPostByPostId(postId, post)

	numberOfComment := service.GetCountCommentByPost(postId)
	userId := decodedToken.ID

	post.Like += userId + ","
	service.UpdateColumnLike(post)
	isAdmin := adminSchema.ID == userId

	return c.JSON(utils.CreatePostResponse(post, decodedToken.Username, decodedToken.ID, isAdmin, numberOfComment))
}
