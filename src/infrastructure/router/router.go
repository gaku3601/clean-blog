package router

import (
	"github.com/gaku3601/clean-blog/src/infrastructure/database"
	"github.com/gaku3601/clean-blog/src/interfaces/controller"
	gin "github.com/gin-gonic/gin"
)

// Router ルーター
var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controller.NewUserController(database.NewSQLHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.POST("/signin", func(c *gin.Context) { userController.SignIn(c) })

	Router = router
}
