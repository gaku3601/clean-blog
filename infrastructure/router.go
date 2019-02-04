package infrastructure

import (
	"github.com/gaku3601/clean-blog/interfaces"
	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := interfaces.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.POST("/signin", func(c *gin.Context) { userController.SignIn(c) })

	Router = router
}
