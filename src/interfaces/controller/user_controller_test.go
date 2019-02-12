package controller

import (
	"testing"

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

func (s *testSqlHandler) InsertUser(email string, password string) (id int, err error) {
	return
}

func (s *testSqlHandler) FetchUserID(email string) (id int, err error) {
	id = 1
	err = nil
	return
}
func (s *testSqlHandler) UpdateUserPassword(id int, hashPassword string) (err error) {
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
