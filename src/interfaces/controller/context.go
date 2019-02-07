package controller

type Context interface {
	UserParams() (email string, password string)
	JSON(int, interface{})
}
