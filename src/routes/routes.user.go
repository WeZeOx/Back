package routes

import (
	"Forum-Back-End/src/controller"
	"Forum-Back-End/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func UsersRouters(router fiber.Router) {
	router.Post("/signup", middleware.CheckFieldCreateUser, controller.CreateUser)
	router.Post("/signin", middleware.CheckFieldLogin, controller.LoginUser)
	router.Get("/all", controller.GetUsers)
	router.Get("/:id", controller.GetUser)
}
