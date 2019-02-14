package router

import (
	"github.com/gaku3601/clean-blog/src/infrastructure/database"
	"github.com/gaku3601/clean-blog/src/infrastructure/mail"
	"github.com/gaku3601/clean-blog/src/interfaces/controller"
	"github.com/gaku3601/clean-blog/src/usecase"
	gin "github.com/gin-gonic/gin"
)

type Context struct{ gin *gin.Context }

func (c *Context) UserParams() (email string, password string) {
	type Json struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var j Json
	c.gin.BindJSON(&j)
	email = j.Email
	password = j.Password
	return
}
func (c *Context) JSON(status int, content interface{}) {
	c.gin.JSON(status, content)
}

func Start() {
	router := gin.Default()

	c := usecase.NewUserUsecase(database.NewSQLHandler(), mail.NewMailHandler())
	userController := controller.NewUserController(c)

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
