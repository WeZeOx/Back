package controller

import (
	"Forum-Back-End/src/dto"
	"Forum-Back-End/src/service"
	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	comment := c.Locals("comment").(dto.ContentCommentCreator)
	token := c.Locals("decodedToken").(*dto.Claims)
	service.CreateComment(comment)
	return c.JSON(fiber.Map{"comment": comment, "isAdmin": token.IsAdmin})
}
