package controller

import (
	"github.com/gaku3601/clean-blog/src/usecase"
)

type UserController struct {
	usecase.UserInterface
}

func NewUserController(user usecase.UserInterface) *UserController {
	return &UserController{
		user,
	}
}

// Create ユーザ作成処置
func (ctrl *UserController) Create(c Context) {
	email := c.EmailParam()
	password := c.PasswordParam()
	err := ctrl.AddUser(email, password)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, nil)
}

// SignIn ログイン処理
func (ctrl *UserController) SignIn(c Context) {
	email := c.EmailParam()
	password := c.PasswordParam()
	token, err := ctrl.GetAccessToken(email, password)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, token)
}

// SendValidEmail email有効化メールを再送信します。
func (ctrl *UserController) SendValidEmail(c Context) {
	email := c.EmailParam()
	err := ctrl.ReSendConfirmValidEmail(email)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "success")
}

//ChangePassword Userのpasswordを変更します
func (ctrl *UserController) ChangePassword(c Context) {
	id := c.IDParam()
	password := c.PasswordParam()
	newPassword := c.NewPasswordParam()
	err := ctrl.ChangeUserPassword(id, password, newPassword)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "success")
}

// ValidEmail tokenを検証し、emailを有効化します。
func (ctrl *UserController) ValidEmail(c Context) {
	token := c.EmailTokenParam()
	err := ctrl.ActivationEmail(token)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "success")
}

// ValidPassword パスワード認証を有効化します。
func (ctrl *UserController) ValidPassword(c Context) {
	id := c.IDParam()
	newPassword := c.NewPasswordParam()
	err := ctrl.ActivationPassword(id, newPassword)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "success")
}
