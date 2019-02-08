package domain

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("Userが生成された際", t, func() {
		Convey("validationが効くか", func() {
			_, err := NewUser("", "password")
			So(err, ShouldNotBeNil)
		})
	})

}

func TestValidation(t *testing.T) {
	Convey("validation check", t, func() {
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

func TestFetchJWT(t *testing.T) {
	Convey("FetchJWT()でtokenが返却されることを確認する", t, func() {
		u := &User{Email: "ex@example.com"}
		t, _ := u.CreateJWT(1)

		token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar"), nil
		})

		Convey("emailが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["email"], ShouldEqual, "ex@example.com")
		})
		Convey("idが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["id"], ShouldEqual, 1)
		})
		Convey("expが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["exp"], ShouldNotBeNil)
		})
		Convey("iatが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["iat"], ShouldNotBeNil)
		})
	})
}
