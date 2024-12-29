package repository

import (
	"book-rent-api/model/domain"
	"context"
)

type BookRepository interface {
	Save(ctx context.Context, Book domain.Book) domain.Book
	Update(ctx context.Context, Book domain.Book) (domain.Book, error)
	Delete(ctx context.Context, Book domain.Book) error
	FindById(ctx context.Context, BookId string) (domain.Book, error)
	FindAll(ctx context.Context) []domain.Book
}