package database

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStore(t *testing.T) {
	Convey("Store()のtest", t, func() {
		Convey("登録に成功した場合、err=nilが格納されていること", func() {
			u := &UserRepository{&testSqlHandler{}}
			err := u.Store("ex@example.com", "test")
			So(err, ShouldBeNil)
		})
	})
}

type testSqlHandler struct{}

func (t *testSqlHandler) InsertUser(email string, password string) (err error) {

	return nil
}
