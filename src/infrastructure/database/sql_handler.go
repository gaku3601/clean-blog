package database

import (
	"database/sql"
	"os"

	"github.com/gaku3601/clean-blog/src/interfaces/database"
	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// SQLHandler ハンドラー
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler : SQLHandlerを返却します
func NewSQLHandler() database.SQLHandler {
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

// InsertUser ユーザ内容を格納します。
func (handler *SQLHandler) InsertUser(email string, password string) (id int, err error) {
	err = handler.Conn.QueryRow("Insert Into users (email, password) values ($1, $2) RETURNING id;", email, password).Scan(&id)
	return
}

func (handler *SQLHandler) FetchUserID(email string) (id int, err error) {
	// TODO: 実装する
	return 0, nil
}
