package service

import (
	"book-rent-api/helper"
	"book-rent-api/model/domain"
	"book-rent-api/model/web"
	"book-rent-api/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	Validate *validator.Validate
	userRepository	repository.UserRepository
}

func NewUserService(validate *validator.Validate, userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		Validate:       validate,
		userRepository: userRepository,
	}
}

func(service *UserServiceImpl) Create(ctx context.Context, request web.CreateUserRequest) (web.UserResponse, error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return web.UserResponse{}, errValidate
	}

	user := domain.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Address:     request.Address,
	}

	user = service.userRepository.Save(ctx, user)

	return helper.ToUserResponse(user),nil

}

func(service *UserServiceImpl) CreateRent(ctx context.Context, request *web.CreateRentRequest) web.RentResponse {
	user := service.userRepository.SaveRent(ctx,request.UserId, request.BookIds)

	return helper.ToRentResponse(user)
}

func(service *UserServiceImpl) Update(ctx context.Context, request *web.UpdateUserRequest) (web.UserResponse, error) {
	user := domain.User{
		ID:          uuid.MustParse(request.ID),
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Address:     request.Address,
	}

	user, err := service.userRepository.Update(ctx,user)

	return helper.ToUserResponse(user), err
}
func(service *UserServiceImpl) Delete(ctx context.Context, requestId string) error {
	user, err := service.userRepository.FindById(ctx, requestId)

	if err != nil {
		panic(err)
	}

	service.userRepository.Delete(ctx, user)

	return err
}
func(service *UserServiceImpl) FindById(ctx context.Context, requestId string) web.UserResponse {
	user, err := service.userRepository.FindById(ctx, requestId)

	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}
func(service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	users := service.userRepository.FindAll(ctx)

	return helper.ToUserResponses(users)
}