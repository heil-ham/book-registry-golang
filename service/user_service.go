package service

import (
	"book-rent-api/model/web"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request web.CreateUserRequest) (web.UserResponse, error)
	CreateRent(ctx context.Context, request *web.CreateRentRequest) web.RentResponse
	Update(ctx context.Context, request *web.UpdateUserRequest) (web.UserResponse, error)
	Delete(ctx context.Context, requestId string) error
	FindById(ctx context.Context, requestId string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}