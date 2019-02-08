package usecase

// UserRepository interface
type UserRepository interface {
	StoreUser(email string, password string) (id int, err error)
	CheckExistUser(email string) (id int, err error)
	CheckCertificationUser(email string, password string) (id int, err error)
	UpdateValidEmail(email string) error
	CreateSocialProfile(servise string, userID int, uid string) error
	CheckExistSocialProfile(servise string, uid string) (userID int, err error)
}
