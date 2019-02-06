package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID           int
	Email        string
	Password     string
	HashPassword string
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
