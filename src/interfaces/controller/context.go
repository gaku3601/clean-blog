package controller

type Context interface {
	EmailParam() (email string)
	PasswordParam() (password string)
	JSON(int, interface{})
}
