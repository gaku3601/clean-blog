package domain

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int
	Email      string
	Password   string
	ValidEmail bool
}

func NewUser(email string) (*User, error) {
	u := &User{Email: email}
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

func (u *User) CreateValidEmailToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now(),
	})
	tokenstring, err := token.SignedString([]byte("foobar2"))

	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func (u *User) CreateHashPassword(password string) (hashPassword string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPassword = string(hash)
	return
}
