package usecase

import (
	"github.com/gaku3601/clean-blog/src/domain"
)

// UserUsecase ユースケースstruct
type UserUsecase struct {
	UserRepository
}

// Add ユーザを追加します。
func (u *UserUsecase) AddUser(email string, password string) error {
	d, err := domain.NewUser(email)
	if err != nil {
		return err
	}
	h := d.CreateHashPassword(password)
	_, err = u.StoreUser(d.Email, h)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) FetchJWT(email string, password string) (string, error) {
	id, err := u.CheckCertificationUser(email, password)
	if err != nil {
		return "", err
	}
	d, err := domain.NewUser(email)
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

func (u *UserUsecase) CertificationSocialProfile(servise ServiseEnum, email string, uid string) (token string, err error) {
	userID, err := u.CheckExistSocialProfile(string(servise), uid)
	if err != nil && err.Error() == "No Data" {
		userID, err := u.CheckExistUser(email)
		if err != nil && err.Error() == "No Data" {
			userID, err := u.StoreNonPasswordUser(email)
			if err != nil {
				return "", err
			}
			u.StoreSocialProfile(string(servise), userID, uid)
			return "", nil
		} else if err != nil {
			return "", err
		}
		err = u.StoreSocialProfile(string(servise), userID, uid)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	d, err := domain.NewUser(email)
	if err != nil {
		return "", err
	}
	token, err = d.CreateJWT(userID)
	if err != nil {
		return "", err
	}
	return token, err
}
