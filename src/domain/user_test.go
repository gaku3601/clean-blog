package domain

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("Userが生成された際", t, func() {
		u := NewUser(1, "email", "password")
		Convey("HashPasswordも格納されるか", func() {
			So(len(u.HashPassword), ShouldEqual, 60)
		})
		Convey("JWTも格納されるか", func() {
			So(len(u.JWT), ShouldNotBeEmpty)
		})
	})

}

func TestCreateHashPassword(t *testing.T) {
	Convey("hash化されているか検証する", t, func() {
		u := &User{}
		u.createHashPassword()
		So(len(u.HashPassword), ShouldEqual, 60)
	})
}

func TestCreateJWT(t *testing.T) {
	Convey("createJWT()でtokenがセットされることを確認する", t, func() {
		u := &User{ID: 1, Email: "ex@example.com"}
		u.createJWT()
		token, _ := jwt.Parse(u.JWT, func(token *jwt.Token) (interface{}, error) {
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
