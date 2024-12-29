package web

type RentResponse struct {
	ID		string	`json:"user_id"`
	Name	string	`json:"name"`
	Books	[]BookRented	`json:"books"`
}

type BookRented struct {
	ID string `json:"book_id"`
	Title string	`json:"title"`
}