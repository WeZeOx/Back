package middleware

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func CheckFieldCreateComment(c *fiber.Ctx) error {
	var checkFieldPostArray = []string{"user_id", "post_id", "content_comment"}
	var comment dto.ContentCommentCreator
	err := c.BodyParser(&comment)

	if (err != nil) ||
		!utils.CheckFieldComment(comment, checkFieldPostArray) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.State{
			Message: "Bad Fields",
			Auth:    false,
		})
	}

	comment.CommentId = uuid.New().String()
	comment.CreatedAt = time.Now()

	c.Locals("comment", comment)
	return c.Next()
}
