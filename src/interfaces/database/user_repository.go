package database

import "github.com/gaku3601/clean-blog/src/domain"

type UserRepository struct {
	SQLHandler
}

// TODO: passwordをhashPasswordへ引数名を変更する
func (repo *UserRepository) StoreUser(email string, password string) (id int, err error) {
	id, err = repo.InsertUser(email, password)
	return
}

func (repo *UserRepository) UpdateUserPassword(id int, hashPassword string) (err error) {
	//TODO: あとで実装する
	return
}

func (repo *UserRepository) GetUser(id int) (user *domain.User, err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserRepository) StoreNonPasswordUser(email string) (id int, err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserRepository) CheckExistUser(email string) (id int, err error) {
	id, err = repo.FetchUserID(email)
	return
}
func (repo *UserRepository) CheckCertificationUser(email string, password string) (id int, err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserRepository) UpdateValidEmail(id int) (err error) {
	// TODO: あとで実装する
	return nil
}

func (repo *UserRepository) StoreSocialProfile(servise string, userID int, uid string) error {
	// TODO: あとで実装する
	return nil
}

func (repo *UserRepository) CheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	// TODO: あとで実装する
	return
}
