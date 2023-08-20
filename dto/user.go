package dto

type UserDTO struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
