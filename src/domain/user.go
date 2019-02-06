package domain

import (
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

type Auth struct {
	Token string
}

func NewUser(password string) *User {
	u := &User{Password: password}
	u.HashPassword = u.createHash(password)
	return u
}

func (u *User) createHash(str string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash)
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
