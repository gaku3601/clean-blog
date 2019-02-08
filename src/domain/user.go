package domain

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID         int
	Email      string
	Password   string
	ValidEmail bool
}

func NewUser(email string, password string) (*User, error) {
	u := &User{Email: email, Password: password}
	err := u.validation()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) validation() error {
	if u.Email == "" {
		return errors.New("Emailを格納してください。")
	}
	if u.Password == "" {
		return errors.New("Passwordを格納してください。")
	}
	return nil
}

func (u *User) CreateJWT(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    id,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now(),
	})
	tokenstring, err := token.SignedString([]byte("foobar"))

	if err != nil {
		return "", err
	}
	return tokenstring, nil
}
