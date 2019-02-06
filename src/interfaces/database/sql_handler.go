package database

type SqlHandler interface {
	InsertUser(email string, password string) error
}
