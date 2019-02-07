package controller

type Context interface {
	ParamsCreate() (email string, password string)
	JSON(int, interface{})
}
