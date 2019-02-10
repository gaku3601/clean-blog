package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/clean-blog/src/domain"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
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
	Convey("ActivationEmail()", t, func() {
		Convey("改ざんされたtokenが渡った場合、errorが返却されること", func() {
			err := u.ActivationEmail("token")
			So(err, ShouldNotBeNil)
		})
		Convey("ActivationEmail()で更新処理に成功した時、nilが返却されること", func() {
			t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
				"id":  12,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
				"iat": time.Now(),
			})
			token, _ := t.SignedString([]byte("foobar2"))
			err := u.ActivationEmail(token)
			So(err, ShouldBeNil)
		})
	})
	Convey("GetAccessToken()", t, func() {
		Convey("Userが存在していない場合、errを返却する", func() {
			_, err := u.GetAccessToken("ng@mail", "ngpass")
			So(err, ShouldNotBeNil)
		})
		Convey("Userが存在している場合、tokenが返却されること", func() {
			token, _ := u.GetAccessToken("ok@mail", "okpass")
			So(token, ShouldNotBeEmpty)
		})
	})
	Convey("ConfirmValidAccessToken()", t, func() {
		Convey("有効なtokenであれば、idが返却されるか", func() {
			t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
				"id":  8,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
				"iat": time.Now(),
			})
			token, _ := t.SignedString([]byte("foobar"))
			id, _ := u.ConfirmValidAccessToken(token)
			So(id, ShouldEqual, 8)
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
	Convey("ChangeUserPassword()", t, func() {
		Convey("現在のpasswordが間違っている場合、errorが返却されること", func() {
			err := u.ChangeUserPassword(1, "ngpassword", "nextPassword")
			So(err, ShouldNotBeNil)
		})
		Convey("現在のpasswordがあっている場合、errorが返却されないこと", func() {
			err := u.ChangeUserPassword(2, "okpassword", "nextPassword")
			So(err, ShouldBeNil)
		})
		Convey("現在のpasswordがあっている場合、更新処理が行われること", func() {
			err := u.ChangeUserPassword(2, "okpassword", "nextPassword")
			So(err, ShouldBeNil)
		})
	})
}

type testRepo struct{}
type testMail struct{}

func (r *testRepo) StoreUser(email string, password string) (id int, err error) {
	return 0, nil
}
func (r *testRepo) GetUser(id int) (user *domain.User, err error) {
	if id == 1 {
		return &domain.User{Password: "ng"}, nil
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("okpassword"), bcrypt.DefaultCost)
	hashPassword := string(hash)
	return &domain.User{Password: hashPassword}, nil
}
func (r *testRepo) UpdateUserPassword(id int, hashPassword string) (err error) {
	return
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

func (r *testRepo) UpdateValidEmail(id int) error {
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
