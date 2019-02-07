package controller

import (
	"github.com/gaku3601/clean-blog/src/domain"
	"github.com/gaku3601/clean-blog/src/interfaces/database"
	"github.com/gaku3601/clean-blog/src/usecase"
)

type UserController struct {
	*usecase.UserUsecase
}

func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		&usecase.UserUsecase{
			&database.UserRepository{
				sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	email, password := c.UserParams()
	err := controller.Add(email, password)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, nil)
}

func (controller *UserController) SignIn(c Context) {
	// TODO: ドメインへのアクセスをやめたい
	u := domain.User{}
	c.JSON(200, u.JWT)
}
