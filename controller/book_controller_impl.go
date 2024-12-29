package controller

import (
	"book-rent-api/model/web"
	"book-rent-api/service"

	"github.com/gofiber/fiber/v2"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(service service.BookService) *BookControllerImpl {
	return &BookControllerImpl{
		BookService: service,
	}
}

func (controller *BookControllerImpl) Create(ctx *fiber.Ctx) error {
	createBookRequest := web.CreateBookRequest{}

	err := ctx.BodyParser(&createBookRequest)

	if err != nil {
		return err
	}

	response := controller.BookService.Create(ctx.Context(), &createBookRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	err = ctx.JSON(webResponse)

	return err
}

func (controller *BookControllerImpl) Update(ctx *fiber.Ctx) error {
	request := web.UpdateBookRequest{}

	err := ctx.BodyParser(&request)

	if err != nil {
		return err
	}

	response, err := controller.BookService.Update(ctx.Context(), &request)

	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	err = ctx.JSON(webResponse)

	return err
}

func (controller *BookControllerImpl) Delete(ctx *fiber.Ctx) error {
	requestId := ctx.Params("bookId")
	err := controller.BookService.Delete(ctx.Context(), requestId)

	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	err = ctx.JSON(webResponse)

	return err
}

func (controller *BookControllerImpl) FindById(ctx *fiber.Ctx) error {
	requestId := ctx.Params("bookId")
	response := controller.BookService.FindById(ctx.Context(), requestId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	err := ctx.JSON(webResponse)

	return err
}

func (controller *BookControllerImpl) FindAll(ctx *fiber.Ctx) error {
	books := controller.BookService.FindAll(ctx.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   books,
	}

	err  := ctx.JSON(webResponse)

	return err
}