package infrastructure

import (
	"database/sql"

	"github.com/gaku3601/clean-blog/interfaces"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() interfaces.SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

type SqlResult struct {
	Result sql.Result
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (interfaces.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}
