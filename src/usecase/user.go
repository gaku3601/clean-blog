package usecase

import (
	"github.com/gaku3601/clean-blog/src/domain"
)

// UserUsecase ユースケースstruct
type UserUsecase struct {
	UserRepository
	UserMail
}

// Add ユーザを追加します。
func (u *UserUsecase) AddUser(email string, password string) error {
	d := new(domain.User)
	h := d.CreateHashPassword(password)
	id, err := u.StoreUser(d.Email, h)
	if err != nil {
		return err
	}
	token := d.CreateValidEmailToken(id)
	u.SendConfirmValidEmail(email, token)
	return nil
}

func (u *UserUsecase) GetAccessToken(email string, password string) (string, error) {
	id, err := u.CheckCertificationUser(email, password)
	if err != nil {
		return "", err
	}
	d := new(domain.User)
	token := d.CreateAccessToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserUsecase) ConfirmValidAccessToken(accessToken string) (id int, err error) {
	d := new(domain.User)
	id, err = d.CheckAccessToken(accessToken)
	return
}

func (u *UserUsecase) ActivationEmail(validToken string) error {
	d := new(domain.User)
	id, err := d.CheckValidEmailToken(validToken)
	if err != nil {
		return err
	}
	err = u.UpdateValidEmail(id)
	return err
}

type ServiseEnum string

const (
	google ServiseEnum = "google"
)

func (u *UserUsecase) CertificationSocialProfile(servise ServiseEnum, email string, uid string) (token string, err error) {
	d := new(domain.User)
	userID, err := u.CheckExistSocialProfile(string(servise), uid)
	if err != nil && err.Error() == "No Data" {
		userID, err := u.CheckExistUser(email)
		if err != nil && err.Error() == "No Data" {
			userID, err := u.StoreNonPasswordUser(email)
			if err != nil {
				return "", err
			}
			u.StoreSocialProfile(string(servise), userID, uid)
			token := d.CreateValidEmailToken(userID)
			u.SendConfirmValidEmail(email, token)
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
	token = d.CreateAccessToken(userID)
	return
}
