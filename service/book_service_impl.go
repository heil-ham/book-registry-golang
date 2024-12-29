package service

import (
	"book-rent-api/helper"
	"book-rent-api/model/domain"
	"book-rent-api/model/web"
	"book-rent-api/repository"
	"context"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	Repository repository.BookRepository
	Validator *validator.Validate
}

func NewBookService(validator *validator.Validate, repository repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{
		Repository: repository,
		Validator:  validator,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request *web.CreateBookRequest) web.BookResponse {
	service.Validator.Struct(request)
	
	book := domain.Book{
		Title:        request.Title,
	}
	
	service.Repository.Save(ctx, book) 

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request *web.UpdateBookRequest) (web.BookResponse, error) {
	book, err := service.Repository.FindById(ctx, request.ID)

	if err != nil {
		panic(err)
	}

	book.Title = request.Title

	response,err := service.Repository.Update(ctx,book)

	return helper.ToBookResponse(response), err
}
func (service *BookServiceImpl) Delete(ctx context.Context, requestId string) error {
	book, err := service.Repository.FindById(ctx, requestId)

	if err != nil {
		panic(err)
	}

	return service.Repository.Delete(ctx, book)
}
func (service *BookServiceImpl) FindById(ctx context.Context, requestId string) web.BookResponse {
	book, err := service.Repository.FindById(ctx, requestId)

	if err != nil {
		panic(err)
	}

	return helper.ToBookResponse(book)
}
func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	books := service.Repository.FindAll(ctx)

	return helper.ToBookResponses(books)
}