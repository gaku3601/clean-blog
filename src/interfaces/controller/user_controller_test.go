package controller

import (
	"testing"

	"github.com/gaku3601/clean-blog/src/usecase"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSignIn(t *testing.T) {
	Convey("SignIn()処理のテスト", t, func() {
		c := NewUserController(&testUser{})
		con := new(testContext)
		c.SignIn(con)
		Convey("200が返却されること", func() {
			So(con.status, ShouldEqual, 200)
		})
		Convey("jwt tokenが返却されること", func() {
			So(con.content, ShouldNotBeEmpty)
		})
	})
}

type testUser struct{}

func (u *testUser) AddUser(email string, password string) (err error) {
	return
}
func (u *testUser) ReSendConfirmValidEmail(email string) (err error) {
	return
}
func (u *testUser) ChangeUserPassword(id int, password string, nextPassword string) (err error) {
	return
}
func (u *testUser) GetAccessToken(email string, password string) (token string, err error) {
	return "token", nil
}
func (u *testUser) ActivationEmail(validToken string) (err error) {
	return
}
func (u *testUser) ActivationPassword(id int, password string) (err error) {
	return
}
func (u *testUser) ForgotPassword(email string) (err error) {
	return
}
func (u *testUser) ProcessForgotPassword(token string, newPassword string) (err error) {
	return
}
func (u *testUser) CertificationSocialProfile(service usecase.ServiceEnum, email string, uid string) (token string, err error) {
	return
}

type testContext struct {
	status  int
	content interface{}
}

func (t *testContext) EmailParam() (email string)             { return "test" }
func (t *testContext) PasswordParam() (password string)       { return }
func (t *testContext) NewPasswordParam() (newPassword string) { return }
func (t *testContext) IDParam() (id int)                      { return }
func (t *testContext) JSON(status int, content interface{}) {
	t.status = status
	t.content = content
}
