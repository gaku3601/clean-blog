package usecase

import "github.com/gaku3601/clean-blog/src/domain"

// UserRepository interface
type UserRepository interface {
	StoreUser(email string, hashPassword string) (id int, err error)
	GetUserByID(id int) (user *domain.User, err error)
	GetUserByEmail(email string) (user *domain.User, err error)
	StoreNonPasswordUser(email string) (id int, err error)
	CheckExistUser(email string) (id int, err error)
	UpdateValidEmail(id int) error
	StoreSocialProfile(servise string, userID int, uid string) error
	CheckExistSocialProfile(servise string, uid string) (userID int, err error)
	UpdateUserPassword(id int, hashPassword string) (err error)
	UpdateActivationPassword(id int, hashPassword string) (err error)
}
