package usecase

import (
	"github.com/gaku3601/clean-blog/src/domain"
)

// UserUsecase ユースケースstruct
type UserUsecase struct {
	UserRepository
}

// Add ユーザを追加します。
func (u *UserUsecase) Add(email string, password string) error {
	d, err := domain.NewUser(email, password)
	if err != nil {
		return err
	}
	h := d.CreateHashPassword()
	err = u.Store(d.Email, h)
	if err != nil {
		return err
	}
	return nil
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

func (u *UserUsecase) ActivationEmail(email string) error {
	err := u.UpdateValidEmail(email)
	return err
}

type ServiseEnum string

const (
	google ServiseEnum = "google"
)

func (u *UserUsecase) CertificationSocialProfile(servise ServiseEnum, email string, uid string) error {
	err := u.CreateSocialProfile(string(servise), email, uid)
	return err
}
