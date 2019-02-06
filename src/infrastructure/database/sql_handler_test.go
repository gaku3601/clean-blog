package database

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSqlHandler(t *testing.T) {
	Convey("DBと接続されていない場合、処理が終了されること", t, func() {
		os.Setenv("DATABASE", "testset")
		defer func() {
			err := recover()
			So(err, ShouldEqual, "DBと接続できませんでした。接続内容を確認してください。")
		}()
		NewSQLHandler()
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

func TestInsertUser(t *testing.T) {
	// DB周りの前処理
	db, err := sql.Open("postgres", fetchDatabaseTestEnv())
	if err != nil {
		panic(err)
	}
	// table作成
	db.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY, email varchar(50) NOT NULL, password  char(60) NOT NULL, UNIQUE(email));")
	Convey("Userが格納可能か検証", t, func() {
		// 関数テスト
		conn, _ := sql.Open("postgres", fetchDatabaseTestEnv())
		s := &SQLHandler{conn}
		s.InsertUser("ex@example.com", "p@ssword")
		// 検証
		var Email string
		var Password string
		fmt.Println(db)
		err := db.QueryRow("select email, password from users where id = 1").Scan(&Email, &Password)
		fmt.Println(err)
		So(Email, ShouldEqual, "ex@example.com")
	})
	//あと処理
	db.Exec("drop schema public cascade;")
	db.Exec("create schema public;")
}

func fetchDatabaseTestEnv() (env string) {
	env = os.Getenv("DATABASE_TEST")
	if env == "" {
		panic("$DATABASE_TEST環境変数を設定してください。")
	}
	return
}
