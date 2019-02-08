package usecase

import (
	"errors"
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

func TestCertificationSocialProfile(t *testing.T) {
	Convey("CertificationSocialProfileのテスト", t, func() {
		r := new(testRepo)
		u := &UserUsecase{r}
		Convey("SocialProfile Tableに既にデータが登録されている場合、JWT tokenを返却する", func() {
			token, _ := u.CertificationSocialProfile(ServiseEnum(google), "ok@example.com", "okuid")
			So(token, ShouldNotBeEmpty)
		})
		Convey("SocialProfile Tableにデータが登録されておらず、User Tableには存在している場合、登録を実施し、JWT tokenを返却する", func() {
			token, _ := u.CertificationSocialProfile(ServiseEnum(google), "ok@example.com", "nguid")
			So(token, ShouldNotBeEmpty)
		})
		Convey("SocialProfile Table、User Table共にデータが存在しない場合、User Table・Social Table共にデータを格納し、JWT tokenを返却しない。", func() {
			token, _ := u.CertificationSocialProfile(ServiseEnum(google), "ng@example.com", "nguid")
			So(token, ShouldBeEmpty)
		})
	})
}

type testRepo struct{}

func (r *testRepo) Store(email string, password string) (id int, err error) {
	return 0, nil
}

func (r *testRepo) CheckExistUser(email string) (userID int, err error) {
	if email == "ok@example.com" {
		userID = 1
		err = nil
		return
	}
	return 0, errors.New("No Data")
}

func (r *testRepo) CheckCertificationUser(email string, password string) (int, error) {
	return 1, nil
}

func (r *testRepo) UpdateValidEmail(email string) error {
	return nil
}

func (r *testRepo) CreateSocialProfile(servise string, userID int, uid string) error {
	return nil
}
func (r *testRepo) CheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	if uid == "okuid" {
		userID = 1
		err = nil
		return
	}
	return 0, errors.New("No Data")
}
