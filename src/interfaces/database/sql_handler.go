package database

type SQLHandler interface {
	InsertUser(email string, password string) error
}
