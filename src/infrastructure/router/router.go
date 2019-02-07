package router

import (
	"github.com/gaku3601/clean-blog/src/infrastructure/database"
	"github.com/gaku3601/clean-blog/src/interfaces/controller"
	gin "github.com/gin-gonic/gin"
)

// Router ルーター
var Router *gin.Engine

type Context struct{}

func (c *Context) ParamsCreate() (email string, password string) { return "", "" }
func (c *Context) JSON(status int, content interface{})          {}

func init() {
	router := gin.Default()

	userController := controller.NewUserController(database.NewSQLHandler())

	router.POST("/users", func(c *gin.Context) {
		con := &Context{}
		userController.Create(con)
	})
	router.POST("/signin", func(c *gin.Context) {
		con := &Context{}
		userController.SignIn(con)
	})

	Router = router
}
