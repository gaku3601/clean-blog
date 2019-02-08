package usecase

import (
	"testing"

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

func TestCreateHashPassword(t *testing.T) {
	Convey("hash化されているか検証する", t, func() {
		hash := createHashPassword("password")
		So(len(hash), ShouldEqual, 60)
	})
}

type testRepo struct{}

func (r *testRepo) Store(email string, password string) error {
	return nil
}

func (r *testRepo) CheckExistUser(email string, password string) (int, error) {
	return 1, nil
}
