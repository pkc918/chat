package dto

type SignUpDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email" validateMsg:"the email is required a legal email"`
	Password string `json:"password" form:"password" validate:"required,min=8" validateMsg:"the password minimum 8 characters required"`
	Avatar   string `json:"avatar" form:"avatar" validate:"required" validateMsg:"avatar must be a legal url"`
}
