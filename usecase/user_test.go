package usecase

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/clean-blog/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateJWT(t *testing.T) {
	Convey("tokenにはemailが格納されていること", t, func() {
		user := domain.User{Email: "ex@example.com"}
		use := new(UserUsecase)
		s := use.CreateJWT(user)
		token, _ := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar"), nil
		})
		So(token.Claims.(jwt.MapClaims)["Email"], ShouldEqual, "ex@example.com")
	})
}
