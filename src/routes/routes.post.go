package routes

import (
	"Forum-Back-End/src/controller"
	"Forum-Back-End/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func PostsRouters(router fiber.Router) {
	router.Get("/all", controller.GetPosts)
	router.Post("/createpost", middleware.CheckToken, middleware.CheckFieldCreatePost, controller.CreatePost)
	router.Patch("/unlike/:postId", middleware.CheckToken, middleware.DecodeToken, controller.UnlikePost)
	router.Patch("/like/:postId", middleware.CheckToken, middleware.DecodeToken, controller.LikePost)
	router.Get("/getpost/:postId", controller.GetSinglePost)
	router.Delete("/deletepost/:postId", middleware.CheckToken, controller.DeletePost)
}
