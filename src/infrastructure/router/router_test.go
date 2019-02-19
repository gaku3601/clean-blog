package router

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gaku3601/clean-blog/src/interfaces/controller"
	gin "github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {
	Convey("ParamsCreateのテスト", t, func() {
		Convey("emailが返却されているかどうか", func() {
			con := &Context{jsonParams: &DecodeJson{Email: "ex@example.com"}}
			email := con.EmailParam()
			So(email, ShouldEqual, "ex@example.com")
		})
	})
	Convey("JSON()のテスト", t, func() {
		Convey("正常にJSONが送信されるかどうか", func() {
			r := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(r)
			con := &Context{gin: c}

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

func TestAuth(t *testing.T) {
	Convey("auth()", t, func() {
		Convey("証明されていないAccessTokenが送信された場合、500エラーが返却されること", func() {
			test := func() *gin.Engine {
				router := gin.Default()
				testFunc := func(controller.Context) {
					return
				}
				router.POST("/", auth(testFunc))

				return router

			}
			ts := httptest.NewServer(test())
			defer ts.Close()
			req, _ := http.NewRequest(
				"POST",
				ts.URL,
				bytes.NewBufferString(""),
			)
			// Content-Type 設定
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", "aaaa")

			client := &http.Client{}
			resp, _ := client.Do(req)
			b, _ := ioutil.ReadAll(resp.Body)
			So(string(b), ShouldEqual, "certification failed.")
		})
	})
}
