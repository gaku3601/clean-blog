package interfaces

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSignIn(t *testing.T) {
	Convey("SignIn()で200が返却されること", t, func() {
		u := new(UserController)
		c := new(testContext)
		u.SignIn(c)
		So(c.status, ShouldEqual, 200)
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
