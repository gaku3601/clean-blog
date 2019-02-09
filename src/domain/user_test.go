package domain

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("FetchJWT()", t, func() {
		u := &User{Email: "ex@example.com"}
		t := u.CreateJWT(1, u.Email)

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
	Convey("CreateValidEmailToken()", t, func() {
		u := &User{}
		t := u.CreateValidEmailToken(1)

		token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar2"), nil
		})

		Convey("tokenにはidが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["id"], ShouldEqual, 1)
		})
		Convey("tokenにはexpが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["exp"], ShouldNotBeNil)
		})
		Convey("tokenにはiatが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["iat"], ShouldNotBeNil)
		})
	})
	Convey("CreateHashPassword()", t, func() {
		u := &User{}
		hash := u.CreateHashPassword("pass")
		So(len(hash), ShouldEqual, 60)
	})
}
