package test

import (
	"book-rent-api/app"
	"book-rent-api/model/web"
	"book-rent-api/repository"
	"book-rent-api/service"
	"context"
	"fmt"

	// "path/filepath"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestEnvFile(t *testing.T) {
	// fmt.Println(filepath.Dir("../config.env")) 
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath("..")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))

}

func TestUserServiceCreate(t *testing.T) {
	validator := validator.New()
	DB := app.NewDB()
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator,userRepository)

	ctx := context.Background()

	userRequest := web.CreateUserRequest{
		Name:        "HEHE",
		PhoneNumber: "+6285997399278",
		Email:       "Sakurgmailcom",
		Address:     "konoha",
	}

	userResponse, err := userService.Create(ctx, userRequest)

	fmt.Println(err)
	fmt.Println(userResponse)
}

func TestUserServiceRentCreate(t *testing.T) {
	validator := new(validator.Validate)
	DB := app.NewDB()
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator,userRepository)

	ctx := context.Background()

	rentRequest := web.CreateRentRequest{
		UserId:  "d4724ed8-a8bc-40b8-8177-474b707c50d3",
		BookIds: []string{
			"6b04447f-92f4-4991-98dc-88333f3b7940",
		},
	}

	rentResponse := userService.CreateRent(ctx, &rentRequest)

	fmt.Println(rentResponse)
}

func TestUserServiceUpdate(t *testing.T) {
	validator := new(validator.Validate)
	DB := app.NewDB()
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator,userRepository)

	ctx := context.Background()

	userRequest := web.UpdateUserRequest{
		ID:          "45192a11-b60e-490c-98d4-043da1b6fc15",
		Name:        "Kakashi",
		PhoneNumber: "+6285997399278",
		Email:       "Kakashi@gmail.com",
		Address:     "konoha",
	}

	userResponse, _ := userService.Update(ctx, &userRequest)

	fmt.Println(userResponse)
}

func TestUserServiceDelete(t *testing.T) {
	validator := new(validator.Validate)
	DB := app.NewDB()
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator,userRepository)

	ctx := context.Background()

	err := userService.Delete(ctx, "b8e52014-62f8-4b1d-afa6-aa439c71fe7a")

	fmt.Println(err)
}

func TestUserServiceFindById(t *testing.T) {
	validator := new(validator.Validate)
	DB := app.NewDB()
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator,userRepository)

	ctx := context.Background()

	user := userService.FindById(ctx,"e19a528b-24f8-463d-8fc8-ad1006585abe")

	fmt.Println(user)
}

func TestUserServiceFindAll(t *testing.T) {
	validator := new(validator.Validate)
	DB := app.NewDB()
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator,userRepository)

	ctx := context.Background()

	users := userService.FindAll(ctx)

	assert.Equal(t, 7, len(users))
}

func TestBookServiceCreate(t *testing.T) {
	validator := validator.New()
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(validator,bookRepository)

	ctx := context.Background()

	bookRequest := web.CreateBookRequest{
		Title:  "Kisah Babat Jawi",
	}

	bookResponse := bookService.Create(ctx, &bookRequest)

	fmt.Println(bookResponse)
}

func TestBookServiceUpdate(t *testing.T) {
	validator := validator.New()
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(validator,bookRepository)

	ctx := context.Background()

	bookRequest := web.UpdateBookRequest{
		ID:    "6781f671-c6ee-4a81-9bb1-0ccfdcb1c5e7",
		Title: "HENTAII",
	}
		
	bookResponse, _ := bookService.Update(ctx, &bookRequest)

	fmt.Println(bookResponse)
}

func TestBookServiceDelete(t *testing.T) {
	validator := validator.New()
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(validator,bookRepository)

	ctx := context.Background()
		
	err := bookService.Delete(ctx, "6781f671-c6ee-4a81-9bb1-0ccfdcb1c5e7")

	assert.Nil(t, err)
}

func TestBookServiceFindById(t *testing.T) {
	validator := validator.New()
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(validator,bookRepository)

	ctx := context.Background()
		
	response := bookService.FindById(ctx, "070d9412-79d8-412c-a994-2b50aab3b33a")

	assert.Equal(t, "One Piece", response.Title)
}

func TestBookServiceFindAll(t *testing.T) {
	validator := validator.New()
	DB := app.NewDB()
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(validator,bookRepository)

	ctx := context.Background()
		
	responses := bookService.FindAll(ctx)

	assert.Equal(t, 6, len(responses))
}

