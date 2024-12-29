package repository

import (
	"book-rent-api/model/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB 	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, user domain.User) domain.User {
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&user).Last(&user).Error

		return err
	})

	return user
}

func (repository *UserRepositoryImpl) SaveRent(ctx context.Context, userId string, bookIds []string) domain.User {
	var err error
	var books []domain.Book
	var user domain.User
	repository.DB.Transaction(func(tx *gorm.DB) error {
		// for _, bookId := range bookIds {
		// 	err = tx.First(&book, "id = ? ",bookId).Error
		// 	if err != nil {
		// 		return err
		// 	}
		// 	books = append(books, book)
		// }
				
		// pengecekan apakah buku tersedia atau tidak
		tx.Where("id IN ?",bookIds).Find(&books)

		err = tx.First(&user, "id = ?", userId).Error

		if err != nil {
			return err
		}

		user.Books = books

		err = tx.Save(&user).Error
		return err
	})

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	var foundedUser domain.User
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Take(&foundedUser, "id = ? ", user.ID.String()).Error

		if err != nil {
			return err
		}

		foundedUser = user
		err = tx.Save(&foundedUser).Error

		return err
	})

	return user, err
}
func (repository *UserRepositoryImpl) Delete(ctx context.Context, user domain.User) error {
	var err error
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&user).Error

		return err
	})
	return err
}
func (repository *UserRepositoryImpl) FindById(ctx context.Context, userId string) (domain.User, error) {
	var user domain.User
	var err error
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.First(&user, "id = ? ", userId).Error

		return err
	})
	return user,err
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context) []domain.User {
	var users []domain.User
	repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&users).Error

		return err
	})
	return users
}