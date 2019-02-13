package controller

import (
	"github.com/gaku3601/clean-blog/src/interfaces/mail"
	"github.com/gaku3601/clean-blog/src/usecase"
)

type UserController struct {
	*usecase.UserUsecase
}

func NewUserController(sqlHandler usecase.UserRepository, mailHandler mail.MailHandler) *UserController {
	return &UserController{
		&usecase.UserUsecase{
			sqlHandler,
			&mail.UserMail{
				mailHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	email, password := c.UserParams()
	err := controller.AddUser(email, password)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, nil)
}

func (controller *UserController) SignIn(c Context) {
	email, password := c.UserParams()
	token, err := controller.GetAccessToken(email, password)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, token)
}
