package database

import (
	"database/sql"
	"os"
	"os/exec"
	"strings"
	"testing"

	_ "github.com/lib/pq"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	conn, _ := sql.Open("postgres", fetchDatabaseTestEnv())
	s := &SQLHandler{conn}
	Convey("StoreUser()", t, func() {
		db := setup()
		defer tearDown()
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
	Convey("GetUser()", t, func() {
		db := setup()
		defer tearDown()
		db.Exec("insert into users (email,password) values ($1,$2)", "ex@mail", "testpass")
		user, _ := s.GetUser(1)
		Convey("IDが格納されているか", func() {
			So(user.ID, ShouldEqual, 1)
		})
		Convey("Emailが格納されているか", func() {
			So(user.Email, ShouldEqual, "ex@mail")
		})
		Convey("Passwordが格納されているか", func() {
			So(strings.TrimSpace(user.Password), ShouldEqual, "testpass")
		})
		Convey("ValidEmailが格納されているか", func() {
			So(user.ValidEmail, ShouldBeFalse)
		})
		Convey("ValidPasswordが格納されているか", func() {
			So(user.ValidPassword, ShouldBeFalse)
		})
	})
	Convey("StoreNonPasswordUser()", t, func() {
		db := setup()
		defer tearDown()
		id, _ := s.StoreNonPasswordUser("ex@example.com")
		// 検証
		var Email string
		var Password string
		var ValidPassword bool
		db.QueryRow("select email, password, valid_password from users where id = 1").Scan(&Email, &Password, &ValidPassword)
		Convey("正常に格納されるか", func() {
			So(Email, ShouldEqual, "ex@example.com")
		})
		Convey("idが返却されているか", func() {
			So(id, ShouldEqual, 1)
		})
		Convey("valid_passwordはfalseとなっているか", func() {
			So(ValidPassword, ShouldEqual, false)
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
