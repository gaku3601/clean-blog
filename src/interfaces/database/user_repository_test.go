package database

import (
	"testing"

	"github.com/gaku3601/clean-blog/src/domain"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStore(t *testing.T) {
	Convey("Store()のtest", t, func() {
		Convey("登録に成功した場合、err=nilが格納されていること", func() {
			d := domain.User{Email: "ex@example.com", Password: "test"}
			u := &UserRepository{&testSqlHandler{}}
			err := u.Store(d)
			So(err, ShouldBeNil)
		})
	})
}

type testSqlHandler struct{}

func (t *testSqlHandler) InsertUser(email string, password string) (err error) {

	return nil
}
