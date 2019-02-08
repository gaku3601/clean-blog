package usecase

// UserRepository interface
type UserRepository interface {
	Store(email string, password string) error
	CheckExistUser(email string, password string) (id int, err error)
	UpdateValidEmail(email string) error
	CreateSocialProfile(servise string, email string, uid string) error
}
