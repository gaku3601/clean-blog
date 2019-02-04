package router

import (
	"github.com/gaku3601/clean-blog/src/infrastructure/database"
	"github.com/gaku3601/clean-blog/src/interfaces"
	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := interfaces.NewUserController(database.NewSqlHandler(""))

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.POST("/signin", func(c *gin.Context) { userController.SignIn(c) })

	Router = router
}
