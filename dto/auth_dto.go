package dto

type SignUpInput struct {
	Email    string `json:"email" bidnd:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LogInInput struct {
	Email    string `json:"email" bidnd:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
