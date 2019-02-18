package mail

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("SendConfirmValidEmail()", t, func() {
		SkipConvey("送信可能か(メール送信してしまうため、skipしている)", func() {
			m := NewMailHandler()
			err := m.SendConfirmValidEmail("pro.gaku@gmail.com", "test")
			So(err, ShouldBeNil)
		})
	})
	Convey("createConfirmValidEmailURL()", t, func() {
		Convey("$FRONTHOST環境変数が取得できない場合、errorが出力されること", func() {
			os.Setenv("FRONTHOST", "")
			_, err := createConfirmValidEmailURL("token")
			So(err.Error(), ShouldEqual, "$FRONTHOST環境変数を設定してください。")
		})
		Convey("URLが返却されること", func() {
			os.Setenv("FRONTHOST", "http://okenv.com/")
			url, _ := createConfirmValidEmailURL("token")
			So(url, ShouldEqual, "http://okenv.com/validemail/token")
		})
	})
	Convey("createForgotPasswordURL()", t, func() {
		Convey("$FRONTHOST環境変数が取得できない場合、errorが出力されること", func() {
			os.Setenv("FRONTHOST", "")
			_, err := createForgotPasswordURL("token")
			So(err.Error(), ShouldEqual, "$FRONTHOST環境変数を設定してください。")
		})
		Convey("URLが返却されること", func() {
			os.Setenv("FRONTHOST", "http://okenv.com/")
			url, _ := createForgotPasswordURL("token")
			So(url, ShouldEqual, "http://okenv.com/forgotpassword/token")
		})
	})
}
