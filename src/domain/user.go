package domain

import (
	"errors"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func NewUser(id int, email string, password string) (*User, error) {
	u := &User{ID: id, Email: email, Password: password}
	err := u.validation()
	if err != nil {
		return nil, err
	}
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
