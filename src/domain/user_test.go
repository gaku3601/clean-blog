package domain

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
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
	Convey("CheckValidEmailToken()", t, func() {
		u := &User{}
		Convey("改ざんされたtokenの場合、errが返却されること", func() {
			_, err := u.CheckValidEmailToken("token")
			So(err, ShouldNotBeNil)
		})
		Convey("改ざんされていないtokenが渡された場合、idが返却されること", func() {
			t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
				"id":  13,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
				"iat": time.Now(),
			})
			token, _ := t.SignedString([]byte("foobar2"))

			id, _ := u.CheckValidEmailToken(token)
			So(id, ShouldEqual, 13)
		})
	})
	Convey("CheckHashPassword()", t, func() {
		Convey("hash化されたpasswordと、通常passwordが一致していない場合、falseを返却する", func() {
			hash, _ := bcrypt.GenerateFromPassword([]byte("ngpass"), bcrypt.DefaultCost)
			hashPassword := string(hash)
			u := new(User)
			isMatch := u.CheckHashPassword("password", hashPassword)
			So(isMatch, ShouldBeFalse)

		})
		Convey("hash化されたpasswordと、通常passwordが一致している場合、trueを返却する", func() {
			hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
			hashPassword := string(hash)
			u := new(User)
			isMatch := u.CheckHashPassword("password", hashPassword)
			So(isMatch, ShouldBeTrue)
		})
	})
}
