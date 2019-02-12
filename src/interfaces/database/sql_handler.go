package database

type SQLHandler interface {
	InsertUser(email string, password string) (id int, err error)
	FetchUserID(email string) (id int, err error)
	UpdateUserPassword(id int, hashPassword string) (err error)
}
