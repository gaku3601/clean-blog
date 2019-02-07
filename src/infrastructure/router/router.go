package router

import (
	"github.com/gaku3601/clean-blog/src/infrastructure/database"
	"github.com/gaku3601/clean-blog/src/interfaces/controller"
	gin "github.com/gin-gonic/gin"
)

type Context struct{ *gin.Context }

func (c *Context) ParamsCreate() (email string, password string) {
	type Json struct {
		Email string `json:"email" binding:"required"`
	}
	var j Json
	c.BindJSON(&j)
	email = j.Email
	return
}
func (c *Context) JSON(status int, content interface{}) {}

func Start() {
	router := gin.Default()

	userController := controller.NewUserController(database.NewSQLHandler())

	router.POST("/users", nomal(userController.Create))
	router.POST("/signin", nomal(userController.SignIn))

	router.Run()
}

func nomal(f func(controller.Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		con := &Context{c}
		f(con)
	}
}
