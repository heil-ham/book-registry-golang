package web

type CreateRentRequest struct {
	UserId	string		`validate:"require" json:"user_id"`
	BookIds	[]string	`validate:"require" json:"book_ids"`
}