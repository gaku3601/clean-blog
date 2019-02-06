package domain

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("Userが生成された際", t, func() {
		Convey("HashPasswordも格納されるか", func() {
			u := NewUser("password")
			So(len(u.HashPassword), ShouldEqual, 60)
		})
	})

}

func TestCreateHash(t *testing.T) {
	Convey("hash化されているか検証する", t, func() {
		u := &User{}
		hash := u.createHash("test")
		So(len(hash), ShouldEqual, 60)
	})
}
