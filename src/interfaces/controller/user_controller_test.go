package controller

import (
	"reflect"
	"testing"

	"github.com/gaku3601/clean-blog/src/domain"
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
		Convey("Auth structが返却されること", func() {
			So(reflect.TypeOf(c.content), ShouldEqual, reflect.TypeOf(new(domain.Auth)))
		})
	})
}

type testContext struct {
	status  int
	content interface{}
}

func (t *testContext) Param(string) string    { return "" }
func (t *testContext) Bind(interface{}) error { return nil }
func (t *testContext) Status(int)             {}
func (t *testContext) JSON(status int, content interface{}) {
	t.status = status
	t.content = content
}
