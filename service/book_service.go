package service

import (
	"book-rent-api/model/web"
	"context"
)

type BookService interface {
	Create(ctx context.Context, request *web.CreateBookRequest) web.BookResponse
	Update(ctx context.Context, request *web.UpdateBookRequest) (web.BookResponse, error)
	Delete(ctx context.Context, requestId string) error
	FindById(ctx context.Context, requestId string) web.BookResponse
	FindAll(ctx context.Context) []web.BookResponse
}