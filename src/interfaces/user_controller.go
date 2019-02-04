package interfaces

import (
	"github.com/gaku3601/clean-blog/src/domain"
	"github.com/gaku3601/clean-blog/src/usecase"
)

type UserController struct {
	use usecase.UserUsecase
}

func NewUserController(sqlHandler SqlHandler) *UserController {
	return &UserController{
		use: usecase.UserUsecase{
			Repo: &UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.use.Add(u)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, nil)
}

func (controller *UserController) SignIn(c Context) {
	u := domain.User{}
	auth := controller.use.CreateJWT(u)
	c.JSON(200, auth)
}
