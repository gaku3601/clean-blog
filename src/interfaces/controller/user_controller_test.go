package controller

import (
	"testing"

	"github.com/gaku3601/clean-blog/src/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSignIn(t *testing.T) {
	Convey("SignIn()処理のテスト", t, func() {
		c := NewUserController(&testSqlHandler{}, &testMailHandler{})
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

type testSqlHandler struct{}
type testMailHandler struct{}

func (s *testSqlHandler) IStoreUser(email string, hashPassword string) (id int, err error) {
	return
}
func (s *testSqlHandler) IGetUser(id int) (user *domain.User, err error) {
	return
}
func (s *testSqlHandler) IStoreNonPasswordUser(email string) (id int, err error) {
	return
}
func (s *testSqlHandler) ICheckExistUser(email string) (id int, err error) {
	return
}
func (s *testSqlHandler) ICheckCertificationUser(email string, password string) (id int, err error) {
	return
}
func (s *testSqlHandler) IUpdateValidEmail(id int) (err error) {
	return
}
func (s *testSqlHandler) IStoreSocialProfile(servise string, userID int, uid string) (err error) {
	return
}
func (s *testSqlHandler) ICheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	return
}
func (s *testSqlHandler) IUpdateUserPassword(id int, hashPassword string) (err error) {
	return
}
func (s *testSqlHandler) IUpdateActivationPassword(id int, hashPassword string) (err error) {
	return
}

type testContext struct {
	status  int
	content interface{}
}

func (t *testContext) UserParams() (email string, password string) { return "test", "aaa" }
func (t *testContext) JSON(status int, content interface{}) {
	t.status = status
	t.content = content
}
