package web

type UserResponse struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"` 
	PhoneNumber	string	`json:"phone_number"` 
	Email		string	`json:"email"`
	Address		string	`json:"address"`
}