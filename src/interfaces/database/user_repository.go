package database

import "github.com/gaku3601/clean-blog/src/domain"

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) StoreUser(email string, hashPassword string) (id int, err error) {
	id, err = repo.InsertUser(email, hashPassword)
	return
}

// UpdateUserPassword passwordを新しく設定します。
func (repo *UserRepository) IUpdateUserPassword(id int, hashPassword string) (err error) {
	err = repo.UpdateUserPassword(id, hashPassword)
	return
}

func (repo *UserRepository) GetUser(id int) (user *domain.User, err error) {
	// TODO: あとで実装する
	return
}

// StoreNonPasswordUser Passwordなしでユーザ登録を実施します。
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

// UpdateActivationPassword ValidPasswordをtrueへ変更し、新しくpasswordを設定します。
func (repo *UserRepository) UpdateActivationPassword(id int, hashPassword string) (err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserRepository) StoreSocialProfile(servise string, userID int, uid string) error {
	// TODO: あとで実装する
	return nil
}

func (repo *UserRepository) CheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	// TODO: あとで実装する
	return
}
