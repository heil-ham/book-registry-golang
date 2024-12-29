package controller

import (
	"book-rent-api/model/web"
	"book-rent-api/service"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: service,
	}
}

func (controller *UserControllerImpl) Create(ctx *fiber.Ctx) error {
	createUserRequest := web.CreateUserRequest{}
	ctx.BodyParser(&createUserRequest)

	response, err := controller.UserService.Create(ctx.Context(), createUserRequest)

	if err != nil {
		panic(err)
	}
	
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	err = ctx.JSON(webResponse)

	return err
}

func (controller *UserControllerImpl) CreateRent(ctx *fiber.Ctx) error {
	request := web.CreateRentRequest{}
	ctx.BodyParser(&request)

	response := controller.UserService.CreateRent(ctx.Context(), &request)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	return ctx.JSON(webResponse)
}


func (controller *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	updateUserRequest := web.UpdateUserRequest{}

	err := ctx.BodyParser(&updateUserRequest)

	if err != nil {
		return(err)
	}

	response, err := controller.UserService.Update(ctx.Context(), &updateUserRequest)

	if err != nil {
		return(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	err = ctx.JSON(webResponse)

	return err
}

func (controller *UserControllerImpl) FindById(ctx *fiber.Ctx) error {
	requestId := ctx.Params("userId")
	
	response := controller.UserService.FindById(ctx.Context(), requestId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	err := ctx.JSON(webResponse)

	return err 
}

func (controller *UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	requestId := ctx.Params("userId")
	
	err := controller.UserService.Delete(ctx.Context(), requestId)

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

func (controller *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {
	users := controller.UserService.FindAll(ctx.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	err := ctx.JSON(webResponse)
	
	return err
}
