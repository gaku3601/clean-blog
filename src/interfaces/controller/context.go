package controller

type Context interface {
	IDParam() (id int)
	EmailParam() (email string)
	PasswordParam() (password string)
	NewPasswordParam() (newPassword string)
	EmailTokenParam() (token string)
	ForgotPasswordTokenParam() (token string)
	JSON(int, interface{})
}
