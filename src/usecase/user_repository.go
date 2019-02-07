package usecase

// UserRepository interface
type UserRepository interface {
	Store(email string, password string) error
	CheckExistUser(email string, password string) (id int, err error)
}
