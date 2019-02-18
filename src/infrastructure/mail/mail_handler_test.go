package mail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("SendConfirmValidEmail()", t, func() {
		SkipConvey("送信可能か(メール送信してしまうため、skipしている)", func() {
			m := NewMailHandler()
			err := m.SendConfirmValidEmail("pro.gaku@gmail.com", "test")
			So(err, ShouldBeNil)
		})
	})
}
