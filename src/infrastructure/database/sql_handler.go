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

func (s *SQLHandler) StoreUser(email string, hashPassword string) (id int, err error) {
	err = s.Conn.QueryRow("Insert Into users (email, password, valid_password) values ($1, $2, $3) RETURNING id;", email, hashPassword, true).Scan(&id)
	return
}
func (s *SQLHandler) GetUser(id int) (user *domain.User, err error) {
	user = new(domain.User)
	err = s.Conn.QueryRow("select id, email, password, valid_email, valid_password from users where id = $1;", id).Scan(&user.ID, &user.Email, &user.Password, &user.ValidEmail, &user.ValidPassword)
	return
}
func (s *SQLHandler) StoreNonPasswordUser(email string) (id int, err error) {
	err = s.Conn.QueryRow("Insert Into users (email) values ($1) RETURNING id;", email).Scan(&id)
	return
}
func (s *SQLHandler) CheckExistUser(email string) (id int, err error) {
	return
}
func (s *SQLHandler) CheckCertificationUser(email string, password string) (id int, err error) {
	return
}
func (s *SQLHandler) UpdateValidEmail(id int) (err error) {
	return
}
func (s *SQLHandler) StoreSocialProfile(servise string, userID int, uid string) (err error) {
	return
}
func (s *SQLHandler) CheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	return
}
func (s *SQLHandler) UpdateUserPassword(id int, hashPassword string) (err error) {
	return
}
func (s *SQLHandler) UpdateActivationPassword(id int, hashPassword string) (err error) {
	return
}
