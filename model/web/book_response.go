package web

type BookResponse struct {
	ID				string	`json:"id"`
	Title			string	`json:"title"` 
	UserId			string	`json:"user_id"` 
	Availability	bool	`json:"availability"`
}