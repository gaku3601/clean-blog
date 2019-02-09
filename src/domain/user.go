package domain

import (
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

func (u *User) CreateAccessToken(id int) (token string) {
	t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now(),
	})
	token, err := t.SignedString([]byte("foobar"))

	if err != nil {
		panic(err)
	}
	return
}

func (u *User) CheckAccessToken(accessToken string) (id int, err error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	if err != nil {
		return 0, err
	}
	id = int(token.Claims.(jwt.MapClaims)["id"].(float64))

	return
}

func (u *User) CreateValidEmailToken(id int) (token string) {
	t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now(),
	})
	token, err := t.SignedString([]byte("foobar2"))

	if err != nil {
		panic(err)
	}
	return
}

func (u *User) CreateHashPassword(password string) (hashPassword string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPassword = string(hash)
	return
}
