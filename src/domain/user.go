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

func (u *User) createHash(str string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash)
}
