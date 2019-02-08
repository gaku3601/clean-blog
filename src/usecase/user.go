package usecase

import (
	"github.com/gaku3601/clean-blog/src/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserUsecase ユースケースstruct
type UserUsecase struct {
	UserRepository
}

// Add ユーザを追加します。
func (u *UserUsecase) Add(email string, password string) (err error) {
	h := createHashPassword(password)
	err = u.Store(email, h)
	return
}

func (u *UserUsecase) FetchJWT(email string, password string) (string, error) {
	id, err := u.CheckExistUser(email, password)
	if err != nil {
		return "", err
	}
	d, err := domain.NewUser(email, password)
	if err != nil {
		return "", err
	}
	token, err := d.CreateJWT(id)
	if err != nil {
		return "", err
	}
	return token, nil

}

func createHashPassword(password string) (hashPassword string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPassword = string(hash)
	return
}
