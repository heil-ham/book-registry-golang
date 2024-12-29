package main

import (
	"book-rent-api/app"
	"book-rent-api/controller"
	"book-rent-api/repository"
	"book-rent-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(validate,userRepo)
	userController := controller.NewUserController(userService)
	
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(validate,bookRepo)
	bookController := controller.NewBookController(bookService)

	fiberApp := fiber.New()


	app.RouterUser(fiberApp.Group("/api/users"),userController)
	app.RouterBook(fiberApp.Group("/api/books"),bookController)
	
	fiberApp.Listen("localhost:3000")
}