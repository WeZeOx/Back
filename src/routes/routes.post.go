package routes

import (
	"Forum-Back-End/src/controller"
	"Forum-Back-End/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func PostsRouters(router fiber.Router) {
	router.Get("/all", controller.GetPosts)
	router.Post("/createpost", middleware.CheckToken, middleware.CheckFieldCreatePost, controller.CreatePost)
	router.Post("/likepost", middleware.CheckToken, controller.CreatePost)
}
