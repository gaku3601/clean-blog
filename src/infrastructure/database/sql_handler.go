package database

import (
	"database/sql"
	"os"

	"github.com/gaku3601/clean-blog/src/domain"
	"github.com/gaku3601/clean-blog/src/usecase"

	// PostgreSQL driver
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

// StoreUser Userを格納します。
func (s *SQLHandler) StoreUser(email string, hashPassword string) (id int, err error) {
	err = s.Conn.QueryRow("Insert Into users (email, password, valid_password) values ($1, $2, $3) RETURNING id;", email, hashPassword, true).Scan(&id)
	return
}

// GetUserByID Userを取得しデータを返却します。
func (s *SQLHandler) GetUserByID(id int) (user *domain.User, err error) {
	user = new(domain.User)
	err = s.Conn.QueryRow("select id, email, password, valid_email, valid_password from users where id = $1;", id).Scan(&user.ID, &user.Email, &user.Password, &user.ValidEmail, &user.ValidPassword)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, domain.NoData
	}
	return
}

// GetUserByEmail EmailからUserを検索し返却する。
func (s *SQLHandler) GetUserByEmail(email string) (user *domain.User, err error) {
	user = new(domain.User)
	err = s.Conn.QueryRow("select id, email, password, valid_email, valid_password from users where email = $1;", email).Scan(&user.ID, &user.Email, &user.Password, &user.ValidEmail, &user.ValidPassword)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, domain.NoData
	}
	return
}

// StoreNonPasswordUserAndSocialProfile PasswordなしでUserを格納、ならびにSocialProfileにも格納する。
func (s *SQLHandler) StoreNonPasswordUserAndSocialProfile(email string, service string, uid string) (id int, err error) {
	tx, err := s.Conn.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		// panicがおきたらロールバック
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	err = tx.QueryRow("Insert Into users (email) values ($1) RETURNING id;", email).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	_, err = tx.Exec("Insert Into social_profiles (service,uid,user_id) values ($1,$2,$3);", service, uid, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return
}

// UpdateValidEmail valid_emailをtrueに変更します。
func (s *SQLHandler) UpdateValidEmail(id int) (err error) {
	_, err = s.Conn.Exec("Update users Set valid_email = $2 Where id = $1;", id, true)
	return
}

// StoreSocialProfile SocialProfileテーブルにデータを格納する。
func (s *SQLHandler) StoreSocialProfile(service string, uid string, userID int) (err error) {
	_, err = s.Conn.Exec("Insert Into social_profiles (service, uid, user_id) values ($1, $2, $3);", service, uid, userID)
	return
}

// CheckExistSocialProfile SocialProfileを検索し、データが存在していた場合、userIDを返却する。
func (s *SQLHandler) CheckExistSocialProfile(service string, uid string) (userID int, err error) {
	err = s.Conn.QueryRow("select user_id from social_profiles where service = $1 and uid = $2;", service, uid).Scan(&userID)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return 0, domain.NoData
	}
	return
}

// UpdateUserPassword Userのpasswordを変更します。
func (s *SQLHandler) UpdateUserPassword(id int, hashPassword string) (err error) {
	_, err = s.Conn.Exec("Update users Set password = $2 Where id = $1", id, hashPassword)
	return
}
func (s *SQLHandler) UpdateActivationPassword(id int, hashPassword string) (err error) {
	return
}
