package test

import (
	"book-rent-api/app"
	"book-rent-api/model/domain"
	"book-rent-api/repository"
	"context"
	"testing"
	"time"

	// "github.com/google/uuid"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/book_rent_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()

	if err != nil {
		panic(err)
	}

	sqlDb.SetConnMaxLifetime(30 * time.Minute)
	sqlDb.SetConnMaxIdleTime(5 * time.Minute)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(5)

	return db

}

var db = OpenConnection()

func TestRepositoryUser(t *testing.T) {
	// crate user
	// user := domain.User{
	// 	Name:        "Rizky",
	// 	PhoneNumber: "885899774",
	// 	Email:       "riz@mail",
	// 	Address:     "sidorejo",
	// }
	
	// db.Transaction(func(tx *gorm.DB) error {
	// 	err := tx.Create(&user).Error

	// 	return err
	// })
	
	var user domain.User
	var book domain.Book
	db.Take(&user, "id = ? ", "30c9cd4f-44ce-4e3b-98ae-26600d3a5639")
	db.Take(&book, "id = ? ", "070d9412-79d8-412c-a994-2b50aab3b33a")

	user.Books = append(user.Books, book)

	db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&user).Error

		return err
	})
}

func TestRepositoryUserUpdate(t *testing.T) {
	DB := app.NewDB()
	userRepository := repository.NewUserRepository(DB)

	ctx := context.Background()

	user := domain.User{
		ID:          uuid.MustParse("8632c06b-3c0b-40ab-9958-aa468af1418a"),
		Name:        "coki",
		PhoneNumber: "95389385",
		Email:       "coki@mail.com",
		Address:     "depok",
	}

	user,_  = userRepository.Update(ctx,user)

	assert.Equal(t, "coki", user.Name)
}

func TestRepositorySaveRent(t *testing.T) {
	repository := repository.NewUserRepository(db)

	ctx := context.Background()
	repository.SaveRent(ctx, "679b6ab6-4aef-49d6-81e6-ea8581569c12", []string{"0fa406cf-ced4-465f-bcb0-3e43a5a6583a", "e061640b-94f0-441e-8291-d72c158f70f0"})


}

func TestRepositoryBook(t *testing.T) {
	book := domain.Book{
		Title:        "One Piece",
		Availability: true,
	}

	db.Create(&book)
}

func TestRepositoryBookUpdate(t *testing.T) {
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)

	ctx := context.Background()

	book := domain.Book{
		ID:           uuid.MustParse("6fa8411b-8b81-41c2-97ca-fa583742d098"),
		Title:        "Kamasutra",
	}

	book, _ = bookRepository.Update(ctx,book)

	assert.Equal(t, "Kamasutra", book.Title)
}

func TestRepositoryBookDelete(t *testing.T) {
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)

	ctx := context.Background()

	book := domain.Book{
		ID:           uuid.MustParse("6fa8411b-8b81-41c2-97ca-fa583742d098"),
		Title:        "Kamasutra",
	}

	err := bookRepository.Delete(ctx,book)

	assert.Nil(t, err)
}

func TestRepositoryBookFindId(t *testing.T) {
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)

	ctx := context.Background()

	var book domain.Book

	book, err := bookRepository.FindById(ctx,"faa6db4a-749c-4ddc-8105-592392640968")

	assert.Nil(t, err)
	assert.Equal(t, "Filosofi Teras", book.Title)
}

func TestRepositoryBookFindAll(t *testing.T) {
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)

	ctx := context.Background()

	var books []domain.Book

	books = bookRepository.FindAll(ctx)

	assert.Equal(t, 6, len(books))
}