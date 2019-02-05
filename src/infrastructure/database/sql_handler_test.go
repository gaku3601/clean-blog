package database

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var localDatabaseEnv string

func TestMain(m *testing.M) {
	localDatabaseEnv = os.Getenv("DATABASE")
	code := m.Run()
	os.Exit(code)
}

func TestNewSqlHandler(t *testing.T) {
	Convey("DBと接続されていない場合、処理が終了されること", t, func() {
		os.Setenv("DATABASE", "testset")
		defer func() {
			err := recover()
			So(err, ShouldEqual, "DBと接続できませんでした。接続内容を確認してください。")
		}()
		NewSqlHandler()
	})
}

func TestFetchDatabaseEnv(t *testing.T) {
	Convey("Envが設定されていない場合、panicで終了すること", t, func() {
		os.Setenv("DATABASE", "")
		defer func() {
			err := recover()
			So(err, ShouldEqual, "$DATABASEを環境変数として設定してください。")
		}()
		fetchDatabaseEnv()
	})
}

func TestExecute(t *testing.T) {
	Convey("実行時、errorが発生する場合、適切にerrorが返却されているか。", t, func() {
		os.Setenv("DATABASE", localDatabaseEnv)
		h := NewSqlHandler()
		h.Execute("CREATE TABLE users(id    CHAR(4)    NOT NULL);")
		_, err := h.Execute("CREATE TABLE users(id    CHAR(4)    NOT NULL);")
		So(err.Error(), ShouldEqual, `pq: relation "users" already exists`)

		// すべてのtableをdropする処理
		h.Execute("drop schema public cascade;")
		h.Execute("create schema public;")
	})
}
