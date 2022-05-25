package routes

import (
	"Forum-Back-End/src/controller"
	"github.com/gofiber/fiber/v2"
)

func PostsRouters(router fiber.Router) {
	router.Get("/all", controller.GetPosts)
}
