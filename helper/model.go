package helper

import (
	"book-rent-api/model/domain"
	"book-rent-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Address:     user.Address,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, web.UserResponse{
			ID:          user.ID.String(),
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			Address:     user.Address,
		})
	}

	return userResponses
}

func ToRentResponse(user domain.User) web.RentResponse {
	var booksRented []web.BookRented

	for _, book := range user.Books {
		booksRented = append(booksRented, web.BookRented{
			ID:    book.ID.String(),
			Title: book.Title,
		})
	}

	return web.RentResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Books: booksRented,
	}
}

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		ID:           book.ID.String(),
		Title:        book.Title,
		UserId:       book.UserId,
		Availability: book.Availability,
	}
}

func ToBookResponses(books []domain.Book) []web.BookResponse {
	var bookResponses []web.BookResponse

	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses
}