package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/clean-blog/src/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserUsecase ユースケースstruct
type UserUsecase struct {
	UserRepository
	UserMail
}

func NewUserUsecase(sqlHandler UserRepository, mailHandler UserMail) *UserUsecase {
	u := &UserUsecase{sqlHandler, mailHandler}
	return u
}

// AddUser ユーザを追加します。
func (u *UserUsecase) AddUser(email string, password string) error {
	d := new(domain.User)
	h := u.createHashPassword(password)
	id, err := u.StoreUser(d.Email, h)
	if err != nil {
		return err
	}
	token, err := u.createToken(id, "emailkey")
	if err != nil {
		return err
	}
	u.SendConfirmValidEmail(email, token)
	return nil
}

// ReSendConfirmValidEmail email有効化メールの再送を行います。
func (u *UserUsecase) ReSendConfirmValidEmail(email string) (err error) {
	id, err := u.CheckExistUser(email)
	if err != nil {
		return err
	}
	token, err := u.createToken(id, "emailkey")
	if err != nil {
		return err
	}
	err = u.SendConfirmValidEmail(email, token)
	return
}

// ChangeUserPassword passwordを変更します。
func (u *UserUsecase) ChangeUserPassword(id int, password string, nextPassword string) (err error) {
	user, err := u.GetUser(id)
	if err != nil {
		return err
	}
	isMatch := u.checkHashPassword(password, user.Password)
	if !isMatch {
		return errors.New("Passwords do not match")
	}
	hashPassword := u.createHashPassword(nextPassword)
	err = u.UpdateUserPassword(id, hashPassword)
	return err
}

// GetAccessToken AccessTokenを返却します
func (u *UserUsecase) GetAccessToken(email string, password string) (string, error) {
	id, err := u.CheckCertificationUser(email, password)
	if err != nil {
		return "", err
	}
	token, err := u.createToken(id, "accesskey")
	if err != nil {
		return "", err
	}
	return token, nil
}

// ConfirmValidAccessToken AccessTokenの有効性をチェックし、UserIDを返却します。
func (u *UserUsecase) ConfirmValidAccessToken(accessToken string) (id int, err error) {
	id, err = u.checkToken(accessToken, "accesskey")
	return
}

// ActivationEmail 登録時にメール宛に発行したtokenを検証し、Emailの有効性を確認、更新します。
func (u *UserUsecase) ActivationEmail(validToken string) error {
	id, err := u.checkToken(validToken, "emailkey")
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

// ActivationPassword Password認証を有効化します。
func (u *UserUsecase) ActivationPassword(id int, password string) error {
	user, err := u.GetUser(id)
	if err != nil {
		return err
	}
	if user.ValidPassword {
		return errors.New("passwordが既に有効です。")
	}
	hashPassword := u.createHashPassword(password)
	err = u.UpdateActivationPassword(id, hashPassword)
	return err
}

// ForgotPassword passwordを忘れた際、email宛にpassword再設定URLを発行する。
func (u *UserUsecase) ForgotPassword(email string) (err error) {
	id, err := u.CheckExistUser(email)
	if err != nil {
		return err
	}
	token, err := u.createToken(id, "forgotkey")
	if err != nil {
		return err
	}
	err = u.SendForgotPasswordMail(email, token)
	return
}

// ProcessForgotPassword パスワードを忘れた際に発行したURLから、新しいパスワードを設定します。
func (u *UserUsecase) ProcessForgotPassword(token string, newPassword string) (err error) {
	id, err := u.checkToken(token, "forgotkey")
	if err != nil {
		return err
	}
	hashPassword := u.createHashPassword(newPassword)
	err = u.UpdateActivationPassword(id, hashPassword)
	return
}

// CertificationSocialProfile OpenID認証を行います。
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
			token, err := u.createToken(userID, "emailkey")
			if err != nil {
				return "", err
			}
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
	token, err = u.createToken(userID, "accesskey")
	return
}

// createHashPassword passwordをhash化し返却します。
func (u *UserUsecase) createHashPassword(password string) (hashPassword string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPassword = string(hash)
	return
}

// checkHashPassword hash化されたpasswordとpasswordが一致するか検証します。
func (u *UserUsecase) checkHashPassword(password string, hashPassword string) (isMatch bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// createToken tokenを作成します。
func (u *UserUsecase) createToken(id int, signKey string) (token string, err error) {
	t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now(),
	})
	token, err = t.SignedString([]byte(signKey))

	if err != nil {
		return "", err
	}
	return
}

// checkToken tokenを検証しuserIDを返却します。
func (u *UserUsecase) checkToken(token string, signKey string) (id int, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err != nil {
		return 0, err
	}
	id = int(t.Claims.(jwt.MapClaims)["id"].(float64))

	return
}
