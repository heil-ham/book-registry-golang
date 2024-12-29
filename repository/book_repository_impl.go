package repository

import (
	"book-rent-api/model/domain"
	"context"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		DB: DB,
	}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, Book domain.Book) domain.Book {
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&Book).Last(&Book).Error

		return err
	})

	return Book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, Book domain.Book) (domain.Book, error) {
	var foundedBook domain.Book
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Take(&foundedBook, "id = ? ", Book.ID.String()).Error

		if err != nil {
			return err
		}

		foundedBook = Book
		// err = tx.Save(&foundedBook).Error
		err = tx.Model(&foundedBook).Omit("user_id").Updates(&foundedBook).Error


		return err
	})

	return Book, err
}
func (repository *BookRepositoryImpl) Delete(ctx context.Context, Book domain.Book) error {
	var err error
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&Book).Error

		return err
	})
	return err
}
func (repository *BookRepositoryImpl) FindById(ctx context.Context, BookId string) (domain.Book, error) {
	var Book domain.Book
	var err error
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.First(&Book, "id = ? ", BookId).Error

		return err
	})
	return Book,err
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context) []domain.Book {
	var Books []domain.Book
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&Books).Error

		return err
	})
	return Books
}