package web

type CreateUserRequest struct {
	Name		string	`validate:"required,min=1,max=100" json:"name"` 
	PhoneNumber	string	`validate:"required" json:"phone_number"` 
	Email		string	`validate:"required,email" json:"email"`
	Address		string	`validate:"required" json:"address"`

	// Name		string	 
	// PhoneNumber	string	
	// Email		string	
	// Address		string	
}