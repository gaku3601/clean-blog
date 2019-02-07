package usecase

// UserRepository interface
type UserRepository interface {
	Store(email string, password string) error
}
