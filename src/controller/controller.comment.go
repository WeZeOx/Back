package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func CreateComment(c *fiber.Ctx) error {
	comment := c.Locals("comment").(dto.ContentCommentCreator)
	token := c.Locals("decodedToken").(*dto.JwtClaims)
	comment.UserId = token.ID
	service.CreateComment(comment)
	return c.JSON(fiber.Map{"comment": comment, "isAdmin": token.IsAdmin, "username": token.Username})
}

func GetSinglePostWithComments(c *fiber.Ctx) error {
	var post dto.Post
	postId := c.Params("postId")
	_ = godotenv.Load(".env")
	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	adminSchema := service.GetUserByEmail(ADMIN_EMAIL)
	post = service.GetPostByPostId(postId, post)

	if post.PostID == "" {
		return c.JSON(dto.ResponseState{Message: "Post does not exist", Auth: false, Token: ""})
	} else {
		comments := service.GetPostWithComments(postId)
		var response []fiber.Map

		for _, comment := range comments {
			response = append(response, fiber.Map{
				"comment": comment,
				"admin":   adminSchema.ID == comment.UserId,
			})
		}
		numberOfComment := service.GetCountCommentByPost(postId)
		singlePost := service.FindPost(postId)
		responseSinglePost := utils.CreateUserPostResponse(singlePost, adminSchema.ID == singlePost.UserID, numberOfComment)

		return c.JSON(dto.CommentsWithPost{
			Comments: response,
			Post:     responseSinglePost,
		})
	}
}

func LikeComment(c *fiber.Ctx) error {
	commentId := c.Params("commentId")
	decodedToken := c.Locals("decodedToken").(*dto.JwtClaims)
	comment := service.GetCommentByCommentId(commentId)
	comment.Like += decodedToken.ID + ","
	service.SaveLikeColumn(comment)
	return c.JSON(fiber.Map{"successful": true})
}

func UnlikeComment(c *fiber.Ctx) error {
	commentId := c.Params("commentId")
	decodedToken := c.Locals("decodedToken").(*dto.JwtClaims)
	comment := service.GetCommentByCommentId(commentId)
	userWhoLikeArr := strings.Split(comment.Like, ",")
	newLikeColumn := ""

	for _, id := range userWhoLikeArr {
		if id != decodedToken.ID {
			newLikeColumn += id + ","
		}
	}

	newLikeColumn = newLikeColumn[:len(newLikeColumn)-1]
	comment.Like = newLikeColumn

	service.SaveLikeColumn(comment)

	return c.SendString("")
}
