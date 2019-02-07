package database

type SQLHandler interface {
	InsertUser(email string, password string) error
	FetchUserID(email string, password string) (id int, err error)
}
