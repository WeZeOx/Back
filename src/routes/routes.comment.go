package routes

import (
	"Forum-Back-End/src/controller"
	"Forum-Back-End/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommentsRouters(router fiber.Router) {
	router.Post("/createcomment", middleware.CheckToken, middleware.DecodeToken, middleware.CheckFieldCreateComment, controller.CreateComment)

}
