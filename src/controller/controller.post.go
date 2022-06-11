package controller

import (
	"Forum-Back-End/src/dto"
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
	arrCom := service.GetCountCommentByPost()

	for idx, post := range posts {
		isAdmin := adminSchema.ID == post.Post.UserID
		res = append(res, utils.CreateUserPostResponse(post, isAdmin, arrCom[idx]))
	}
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
	post.PostID = c.Params("postId")
	service.DeletePost(post)

	return c.JSON(fiber.Map{
		"isOk": true,
	})
}

func UnlikePost(c *fiber.Ctx) error {
	var post dto.Post
	postId := c.Params("postId")
	decodedToken := c.Locals("decodedToken").(*dto.Claims)

	post = service.GetPostByPostId(postId, post)
	userId := decodedToken.ID
	newLikeColumn := ""
	userWhoLikeArr := strings.Split(post.Like, ",")

	for _, id := range userWhoLikeArr {
		if id != userId {
			newLikeColumn += id + ","
		}
	}
	newLikeColumn = newLikeColumn[:len(newLikeColumn)-1]
	post.Like = newLikeColumn
	service.UpdateColumnLike(post)

	return c.JSON(post)
}

func LikePost(c *fiber.Ctx) error {
	var post dto.Post
	postId := c.Params("postId")
	decodedToken := c.Locals("decodedToken").(*dto.Claims)
	post = service.GetPostByPostId(postId, post)
	userId := decodedToken.ID
	post.Like += userId + ","
	service.UpdateColumnLike(post)

	return c.JSON(post)
}

func GetSinglePost(c *fiber.Ctx) error {
	var post dto.Post
	postId := c.Params("postId")
	post = service.GetPostByPostId(postId, post)

	if post.PostID == "" {
		return c.JSON(dto.State{
			Message: "Post does not exist",
			Auth:    false,
			Token:   "",
		})
	} else {

		res := service.GetPostWithComments(postId)

		return c.JSON(res)
	}
}
