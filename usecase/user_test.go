package usecase

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/clean-blog/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateJWT(t *testing.T) {
	Convey("CreateJWT()でtokenが返却されたとき", t, func() {
		user := domain.User{ID: 1, Email: "ex@example.com"}
		use := new(UserUsecase)
		s := use.CreateJWT(user)
		token, _ := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar"), nil
		})

		Convey("emailが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["Email"], ShouldEqual, "ex@example.com")
		})
		Convey("idが格納されていること", func() {
			So(token.Claims.(jwt.MapClaims)["ID"], ShouldEqual, 1)
		})

	})
}
