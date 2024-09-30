package dto

type CustomerDTO struct {
	UserDNI  string `json:"user_dni"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
