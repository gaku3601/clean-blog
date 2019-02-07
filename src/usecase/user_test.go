package usecase

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Add()で格納に成功した場合、nilが返却されること", t, func() {
		r := new(testRepo)
		u := &UserUsecase{r}
		err := u.Add("email", "password")

		So(err, ShouldBeNil)
	})
}

func TestFetchJWT(t *testing.T) {
	Convey("FetchJWT()でtokenが返却されることを確認する", t, func() {
		r := new(testRepo)
		u := &UserUsecase{r}
		t, _ := u.FetchJWT("ex@example.com", "password")
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

type testRepo struct{}

func (r *testRepo) Store(email string, password string) error {
	return nil
}

func (r *testRepo) CheckExistUser(email string, password string) (int, error) {
	return 1, nil
}
