package domain

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	Password     string
	HashPassword string
	JWT          string
}

func NewUser(id int, email string, password string) (*User, error) {
	u := &User{ID: id, Email: email, Password: password}
	err := u.validation()
	if err != nil {
		return nil, err
	}
	u.createHashPassword()
	u.createJWT()
	return u, nil
}

func (u *User) validation() error {
	if u.ID == 0 {
		return errors.New("IDを格納してください。")
	}
	if u.Email == "" {
		return errors.New("Emailを格納してください。")
	}
	if u.Password == "" {
		return errors.New("Passwordを格納してください。")
	}
	return nil
}

func (u *User) createHashPassword() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.HashPassword = string(hash)
}

// CreateJWT JWTトークンを作成します。
func (u *User) createJWT() {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now(),
	})
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	u.JWT = tokenstring
}
