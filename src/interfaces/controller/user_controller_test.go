package controller

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSignIn(t *testing.T) {
	Convey("SignIn()処理のテスト", t, func() {
		u := new(UserController)
		c := new(testContext)
		u.SignIn(c)
		Convey("200が返却されること", func() {
			So(c.status, ShouldEqual, 200)
		})
		Convey("jwt tokenが返却されること", func() {
			So(reflect.TypeOf(c.content), ShouldEqual, reflect.TypeOf(""))
		})
	})
}

type testContext struct {
	status  int
	content interface{}
}

func (t *testContext) ParamsCreate() (email string, password string) { return "", "" }
func (t *testContext) JSON(status int, content interface{}) {
	t.status = status
	t.content = content
}
