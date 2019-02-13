package database

import (
	"database/sql"
	"os"

	// PostgreSQL driver
	"github.com/gaku3601/clean-blog/src/domain"
	"github.com/gaku3601/clean-blog/src/usecase"
	_ "github.com/lib/pq"
)

// SQLHandler ハンドラー
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler : SQLHandlerを返却します
func NewSQLHandler() usecase.UserRepository {
	connector := fetchDatabaseEnv()
	conn, _ := sql.Open("postgres", connector)
	err := conn.Ping()
	if err != nil {
		panic("DBと接続できませんでした。接続内容を確認してください。")
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func fetchDatabaseEnv() (d string) {
	d = os.Getenv("DATABASE")
	if d == "" {
		panic("$DATABASEを環境変数として設定してください。")
	}
	return
}

func (s *SQLHandler) IStoreUser(email string, hashPassword string) (id int, err error) {
	return
}
func (s *SQLHandler) IGetUser(id int) (user *domain.User, err error) {
	return
}
func (s *SQLHandler) IStoreNonPasswordUser(email string) (id int, err error) {
	return
}
func (s *SQLHandler) ICheckExistUser(email string) (id int, err error) {
	return
}
func (s *SQLHandler) ICheckCertificationUser(email string, password string) (id int, err error) {
	return
}
func (s *SQLHandler) IUpdateValidEmail(id int) (err error) {
	return
}
func (s *SQLHandler) IStoreSocialProfile(servise string, userID int, uid string) (err error) {
	return
}
func (s *SQLHandler) ICheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	return
}
func (s *SQLHandler) IUpdateUserPassword(id int, hashPassword string) (err error) {
	return
}
func (s *SQLHandler) IUpdateActivationPassword(id int, hashPassword string) (err error) {
	return
}
