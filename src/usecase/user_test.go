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

func TestUpdateValidEmail(t *testing.T) {
	Convey("UpdateValidEmail()で更新処理に成功した時、nilが返却されること", t, func() {
		r := new(testRepo)
		u := &UserUsecase{r}
		err := u.ActivationEmail("ex@example.com")
		So(err, ShouldBeNil)
	})
}

func TestAddSocialProfile(t *testing.T) {
	Convey("AddSocialProfile()で格納に成功した時、nilが返却されること", t, func() {
		r := new(testRepo)
		u := &UserUsecase{r}
		err := u.AddSocialProfile(ServiseEnum(google), "ex@example.com", "10")
		So(err, ShouldBeNil)

	})
}

type testRepo struct{}

func (r *testRepo) Store(email string, password string) error {
	return nil
}

func (r *testRepo) CheckExistUser(email string, password string) (int, error) {
	return 1, nil
}

func (r *testRepo) UpdateValidEmail(email string) error {
	return nil
}

func (r *testRepo) CreateSocialProfile(servise string, email string, uid string) error {
	return nil
}
