package database

import "github.com/gaku3601/clean-blog/src/domain"

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) IStoreUser(email string, hashPassword string) (id int, err error) {
	id, err = repo.InsertUser(email, hashPassword)
	return
}

// UpdateUserPassword passwordを新しく設定します。
func (repo *UserRepository) IUpdateUserPassword(id int, hashPassword string) (err error) {
	err = repo.UpdateUserPassword(id, hashPassword)
	return
}

func (repo *UserRepository) IGetUser(id int) (user *domain.User, err error) {
	// TODO: あとで実装する
	return
}

// StoreNonPasswordUser Passwordなしでユーザ登録を実施します。
func (repo *UserRepository) IStoreNonPasswordUser(email string) (id int, err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserRepository) ICheckExistUser(email string) (id int, err error) {
	id, err = repo.FetchUserID(email)
	return
}
func (repo *UserRepository) ICheckCertificationUser(email string, password string) (id int, err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserRepository) IUpdateValidEmail(id int) (err error) {
	// TODO: あとで実装する
	return nil
}

// UpdateActivationPassword ValidPasswordをtrueへ変更し、新しくpasswordを設定します。
func (repo *UserRepository) IUpdateActivationPassword(id int, hashPassword string) (err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserRepository) IStoreSocialProfile(servise string, userID int, uid string) error {
	// TODO: あとで実装する
	return nil
}

func (repo *UserRepository) ICheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	// TODO: あとで実装する
	return
}
