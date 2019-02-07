package domain

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("Userが生成された際", t, func() {
		Convey("validationが効くか", func() {
			_, err := NewUser(1, "", "password")
			So(err, ShouldNotBeNil)
		})
	})

}

func TestValidation(t *testing.T) {
	Convey("validation check", t, func() {
		Convey("IDが格納されていない時、errorが返却される", func() {
			u := &User{}
			err := u.validation()
			So(err.Error(), ShouldEqual, "IDを格納してください。")
		})
		Convey("Emailが格納されていない時、errorが返却される", func() {
			u := &User{ID: 1}
			err := u.validation()
			So(err.Error(), ShouldEqual, "Emailを格納してください。")
		})
		Convey("Passwordが格納されていない時、errorが返却される", func() {
			u := &User{ID: 1, Email: "aaaa"}
			err := u.validation()
			So(err.Error(), ShouldEqual, "Passwordを格納してください。")
		})
	})
}
