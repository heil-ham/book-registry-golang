package web

type UpdateBookRequest struct {
	ID			string	`validate:"required" json:"id"`
	Title		string	`validate:"required,min=1,max=100" json:"title"`
}