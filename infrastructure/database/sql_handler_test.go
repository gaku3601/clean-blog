package database

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSqlHandler(t *testing.T) {
	Convey("DBと接続されていない場合、処理が終了されること", t, func() {
		defer func() {
			err := recover()
			So(err, ShouldEqual, "DBと接続できませんでした。接続内容を確認してください。")
		}()
		NewSqlHandler("host=127.0.0.1 port=5555 user=root password=password dbname=testdb sslmode=disable")
	})
}
