package database

import (
	"database/sql"
	"os"
	"os/exec"
	"testing"

	_ "github.com/lib/pq"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("StoreUser()", t, func() {
		db := setup()
		defer tearDown()
		conn, _ := sql.Open("postgres", fetchDatabaseTestEnv())
		s := &SQLHandler{conn}
		id, _ := s.StoreUser("ex@example.com", "p@ssword")
		Convey("正常に格納されるか", func() {
			// 検証
			var Email string
			var Password string
			db.QueryRow("select email, password from users where id = 1").Scan(&Email, &Password)
			So(Email, ShouldEqual, "ex@example.com")
		})
		Convey("idが返却されているか", func() {
			So(id, ShouldEqual, 1)
		})
	})
}
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

func fetchDatabaseTestEnv() (env string) {
	env = os.Getenv("DATABASE_TEST")
	if env == "" {
		panic("$DATABASE_TEST環境変数を設定してください。")
	}
	return
}

func setup() *sql.DB {
	db, err := sql.Open("postgres", fetchDatabaseTestEnv())
	if err != nil {
		panic(err)
	}
	exec.Command("goose", "-dir", getMigrationDir(), "postgres", fetchDatabaseTestEnv(), "up").Run()
	return db
}

func tearDown() {
	exec.Command("goose", "-dir", getMigrationDir(), "postgres", fetchDatabaseTestEnv(), "reset").Run()
}

func getMigrationDir() string {
	m := os.Getenv("MIGRATION_DIR")
	if m == "" {
		panic("$MIGRATION_DIRを環境変数として設定してください。")
	}
	return m
}
