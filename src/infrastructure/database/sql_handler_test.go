package database

import (
	"database/sql"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/gaku3601/clean-blog/src/domain"
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
		Convey("valid_passwordはtrueが格納されているか", func() {
			So(ValidPassword, ShouldBeTrue)
		})
	})
	Convey("GetUserByID()", t, func() {
		db := setup()
		defer tearDown()
		db.Exec("insert into users (email,password) values ($1,$2)", "ex@mail", "testpass")
		user, _ := s.GetUserByID(1)
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
		_, err := s.GetUserByID(2)
		Convey("dataが存在しない場合、NoData Errorが返却されるか", func() {
			So(err, ShouldEqual, domain.NoData)
		})
	})
	Convey("GetUserByEmail()", t, func() {
		db := setup()
		defer tearDown()
		db.Exec("insert into users (email,password) values ($1,$2)", "ex@mail", "testpass")
		user, _ := s.GetUserByEmail("ex@mail")
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
		_, err := s.GetUserByEmail("no@mail")
		Convey("dataが存在しない場合、NoData Errorが返却されるか", func() {
			So(err, ShouldEqual, domain.NoData)
		})
	})
	Convey("StoreNonPasswordUserAndSocialProfile()", t, func() {
		db := setup()
		defer tearDown()
		id, _ := s.StoreNonPasswordUserAndSocialProfile("ex@example.com", "google", "okuid")
		// 検証
		var email string
		var validPassword bool
		db.QueryRow("select email, valid_password from users where id = 1").Scan(&email, &validPassword)
		var uid string
		db.QueryRow("select uid from social_profiles where id = 1").Scan(&uid)
		Convey("user tableに正常に格納されるか", func() {
			So(email, ShouldEqual, "ex@example.com")
		})
		Convey("social profile tableに正常に格納されるか", func() {
			So(uid, ShouldEqual, "okuid")
		})
		Convey("idが返却されているか", func() {
			So(id, ShouldEqual, 1)
		})
		Convey("valid_passwordはfalseとなっているか", func() {
			So(validPassword, ShouldEqual, false)
		})
	})
	Convey("UpdateValidEmail()", t, func() {
		db := setup()
		defer tearDown()
		db.Exec("insert into users (email,password) values ($1,$2)", "ex@mail", "testpass")

		s.UpdateValidEmail(1)

		var validEmail bool
		db.QueryRow("select valid_email from users where id = 1").Scan(&validEmail)
		Convey("valid_emailはtrueとなっているか", func() {
			So(validEmail, ShouldBeTrue)
		})
	})
	Convey("StoreSocialProfile()", t, func() {
		db := setup()
		defer tearDown()
		db.Exec("insert into users (email,password) values ($1,$2)", "ex@mail", "testpass")
		s.StoreSocialProfile("google", "okuid", 1)
		// 検証
		var service string
		var uid string
		var userID int
		db.QueryRow("select service,uid,user_id from social_profiles where id = 1").Scan(&service, &uid, &userID)
		Convey("serviseは格納されているか", func() {
			So(service, ShouldEqual, "google")
		})
		Convey("uidは格納されているか", func() {
			So(uid, ShouldEqual, "okuid")
		})
		Convey("user_idは格納されているか", func() {
			So(userID, ShouldEqual, 1)
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
