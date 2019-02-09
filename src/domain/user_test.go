package domain

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("CreateAccessToken()", t, func() {
		u := &User{}
		t := u.CreateAccessToken(1)

		token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar"), nil
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
	Convey("CheckAccessToken()", t, func() {
		u := &User{}
		Convey("改ざんされたtokenの場合、errが返却されること", func() {
			_, err := u.CheckAccessToken("token")
			So(err, ShouldNotBeNil)
		})
		Convey("改ざんされていないtokenが渡された場合、idが返却されること", func() {
			t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
				"id":  12,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
				"iat": time.Now(),
			})
			token, _ := t.SignedString([]byte("foobar"))

			id, _ := u.CheckAccessToken(token)
			So(id, ShouldEqual, 12)
		})
	})
}
