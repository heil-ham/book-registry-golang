package app

import (
	"book-rent-api/controller"

	"github.com/gofiber/fiber/v2"
)

func RouterUser(router fiber.Router, userController controller.UserController) {
	router.Post("/", userController.Create)
	router.Post("/rent", userController.CreateRent)
	router.Get("/:userId", userController.FindById)
	router.Put("/", userController.Update)
	router.Delete("/:userId", userController.Delete)
	router.Get("/", userController.FindAll)
}

func RouterBook(router fiber.Router, bookController controller.BookController) {
	router.Post("/", bookController.Create)
	router.Get("/:bookId", bookController.FindById)
	router.Put("/", bookController.Update)
	router.Delete("/:bookId", bookController.Delete)
	router.Get("/", bookController.FindAll)
}