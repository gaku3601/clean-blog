package domain

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateHash(t *testing.T) {
	Convey("hash化されているか検証する", t, func() {
		u := &User{}
		hash := u.createHash("test")
		So(len(hash), ShouldEqual, 60)
	})
}
