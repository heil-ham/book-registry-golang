package repository

import (
	"book-rent-api/model/domain"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) domain.User
	SaveRent(ctx context.Context, userId string, bookIds []string) domain.User
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
	FindById(ctx context.Context, userId string) (domain.User, error)
	FindAll(ctx context.Context) []domain.User
}