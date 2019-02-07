package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	gin "github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {
	Convey("ParamsCreateのテスト", t, func() {
		Convey("emailが返却されているかどうか", func() {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request, _ = http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"email": "ex@example.com"}`)))
			con := &Context{c}
			email, _ := con.ParamsCreate()
			So(email, ShouldEqual, "ex@example.com")
		})
	})
}
