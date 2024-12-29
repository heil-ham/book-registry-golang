package web

type CreateBookRequest struct {
	Title		string	`validate:"required,min=1,max=100" json:"title"` 
}