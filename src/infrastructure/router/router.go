package router

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/clean-blog/src/infrastructure/database"
	"github.com/gaku3601/clean-blog/src/infrastructure/mail"
	"github.com/gaku3601/clean-blog/src/interfaces/controller"
	"github.com/gaku3601/clean-blog/src/usecase"
	gin "github.com/gin-gonic/gin"
)

type Context struct {
	gin              *gin.Context
	accessTokenParam *DecodeAccessToken
	jsonParams       *DecodeJson
}

func (c *Context) EmailParam() (email string) {
	email = c.jsonParams.Email
	return
}
func (c *Context) PasswordParam() (password string) {
	password = c.jsonParams.Password
	return
}
func (c *Context) NewPasswordParam() (newPassword string) {
	newPassword = c.jsonParams.NewPassword
	return
}
func (c *Context) IDParam() (id int) {
	id = c.accessTokenParam.ID
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
	router.POST("/changepass", auth(userController.ChangePassword))

	router.Run()
}

// nomal 認証なしでアクセスします
func nomal(f func(controller.Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		var j DecodeJson
		c.BindJSON(&j)
		con := &Context{gin: c, jsonParams: &j}
		f(con)
	}
}

// auth 認証ありでアクセスします
func auth(f func(controller.Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		d := DecodeAccessToken{}
		accessToken := c.GetHeader("Authorization")
		_, err := jwt.ParseWithClaims(accessToken, &d, func(token *jwt.Token) (interface{}, error) {
			return []byte("accesskey"), nil // TODO: 環境変数化
		})
		if err != nil {
			c.String(http.StatusBadRequest, "certification failed.")
			return
		}
		var j DecodeJson
		c.BindJSON(&j)
		con := &Context{gin: c, accessTokenParam: &d, jsonParams: &j}
		f(con)
	}
}

type DecodeAccessToken struct {
	ID  int    `json:"id"`
	Iat string `json:"iat"`
	Exp int    `json:"exp"`
	jwt.StandardClaims
}

type DecodeJson struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}
