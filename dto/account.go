package dto

type SignUpDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email" validateMsg:"the email is required a legal email"`
	Password string `json:"password" form:"password" validate:"required,md5" validateMsg:"the password is required a legal MD5 hash"`
	Avatar   string `json:"avatar" form:"avatar" validate:"required" validateMsg:"avatar must be a legal url"`
}

type SignInDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email" validateMsg:"the email is required a legal email"`
	Password string `json:"password" form:"password" validate:"required,md5" validateMsg:"the password is required a legal MD5 hash"`
	Captcha  string `json:"captcha" form:"captcha" validate:"required,min=6" validateMsg:"the captcha is minimum 6 characters required"`
}
