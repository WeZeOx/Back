package routes

import (
	"Forum-Back-End/src/controller"
	"Forum-Back-End/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommentsRouters(router fiber.Router) {
	router.Get("/getpost/:postId", controller.GetSinglePostWithComments)
	router.Post("/createcomment", middleware.CheckToken, middleware.DecodeToken, middleware.CheckFieldCreateComment, controller.CreateComment)
	router.Patch("/like/:commentId", middleware.CheckToken, middleware.DecodeToken, controller.LikeComment)
	router.Patch("/unlike/:commentId", middleware.CheckToken, middleware.DecodeToken, controller.UnlikeComment)
	router.Delete("/deletecomment/:commentId", controller.DeleteComment)
}
