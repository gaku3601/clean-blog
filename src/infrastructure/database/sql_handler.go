package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gaku3601/clean-blog/src/interfaces/database"
)

// SQLHandler ハンドラー
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler : SQLHandlerを返却します
func NewSQLHandler() database.SqlHandler {
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
func (handler *SQLHandler) InsertUser(email string, password string) (err error) {
	_, err = handler.Conn.Exec("Insert Into users (email, password) values ($1, $2);", email, password)
	fmt.Println(err)
	return nil
}
