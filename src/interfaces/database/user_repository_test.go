package database

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStore(t *testing.T) {
	Convey("Store()のtest", t, func() {
		Convey("登録に成功した場合、err=nilが格納されていること", func() {
			u := &UserRepository{&testSqlHandler{}}
			_, err := u.StoreUser("ex@example.com", "test")
			So(err, ShouldBeNil)
		})
	})
	Convey("CheckExistUser()のテスト", t, func() {
		Convey("ユーザが存在している場合、idが返却されること", func() {
			u := &UserRepository{&testSqlHandler{}}
			id, _ := u.CheckExistUser("email")
			So(id, ShouldEqual, 1)
		})
	})
}

type testSqlHandler struct{}

func (t *testSqlHandler) InsertUser(email string, password string) (id int, err error) {

	return 0, nil
}
func (t *testSqlHandler) FetchUserID(email string) (id int, err error) {
	id = 1
	err = nil
	return
}

func (t *testSqlHandler) UpdateUserPassword(id int, hashPassword string) (err error) {
	return
}
