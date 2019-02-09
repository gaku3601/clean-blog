package usecase

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	r := new(testRepo)
	m := new(testMail)
	u := &UserUsecase{r, m}
	Convey("AddUser()", t, func() {
		Convey("AddUser()で格納に成功した場合、nilが返却されること", func() {
			err := u.AddUser("email", "password")

			So(err, ShouldBeNil)
		})
	})
	Convey("UpdateValidEmail()", t, func() {
		Convey("UpdateValidEmail()で更新処理に成功した時、nilが返却されること", func() {
			err := u.ActivationEmail("ex@example.com")
			So(err, ShouldBeNil)
		})
	})
	Convey("FetchAuthToken()", t, func() {
		Convey("Userが存在していない場合、errを返却する", func() {
			_, err := u.FetchAuthToken("ng@mail", "ngpass")
			So(err, ShouldNotBeNil)
		})
		Convey("Userが存在している場合、tokenが返却されること", func() {
			token, _ := u.FetchAuthToken("ok@mail", "okpass")
			So(token, ShouldNotBeEmpty)
		})
	})
	Convey("CertificationSocialProfile()", t, func() {
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
type testMail struct{}

func (r *testRepo) StoreUser(email string, password string) (id int, err error) {
	return 0, nil
}
func (r *testRepo) StoreNonPasswordUser(email string) (id int, err error) {
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
	if email == "ng@mail" {
		return 0, errors.New("errorだよ！")
	}
	return 1, nil
}

func (r *testRepo) UpdateValidEmail(email string) error {
	return nil
}

func (r *testRepo) StoreSocialProfile(servise string, userID int, uid string) error {
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

func (m *testMail) SendConfirmValidEmail(email string, token string) (err error) {
	return
}
