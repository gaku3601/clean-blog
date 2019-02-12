package usecase

import "github.com/gaku3601/clean-blog/src/domain"

// UserRepository interface
type UserRepository interface {
	IStoreUser(email string, hashPassword string) (id int, err error)
	IGetUser(id int) (user *domain.User, err error)
	IStoreNonPasswordUser(email string) (id int, err error)
	ICheckExistUser(email string) (id int, err error)
	ICheckCertificationUser(email string, password string) (id int, err error)
	IUpdateValidEmail(id int) error
	IStoreSocialProfile(servise string, userID int, uid string) error
	ICheckExistSocialProfile(servise string, uid string) (userID int, err error)
	IUpdateUserPassword(id int, hashPassword string) (err error)
	IUpdateActivationPassword(id int, hashPassword string) (err error)
}
