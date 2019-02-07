package router

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
			email, _ := con.UserParams()
			So(email, ShouldEqual, "ex@example.com")
		})
		Convey("passwordが返却されているかどうか", func() {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request, _ = http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"password": "pass"}`)))
			con := &Context{c}
			_, pass := con.UserParams()
			So(pass, ShouldEqual, "pass")
		})
	})
	Convey("JSON()のテスト", t, func() {
		Convey("正常にJSONが送信されるかどうか", func() {
			r := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(r)
			con := &Context{c}

			type Restaurant struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			}
			restaurant := Restaurant{
				Id:   3,
				Name: "サイゼリヤ",
			}
			con.JSON(200, restaurant)

			// 結果を受け取り確認
			body, _ := ioutil.ReadAll(r.Result().Body)
			var restau Restaurant
			json.Unmarshal(body, &restau)
			So(restau.Id, ShouldEqual, 3)
		})
	})
}
