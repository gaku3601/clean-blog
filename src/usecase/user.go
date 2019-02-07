package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// UserUsecase ユースケースstruct
type UserUsecase struct {
	UserRepository
}

// Add ユーザを追加します。
func (u *UserUsecase) Add(email string, password string) (err error) {
	h := createHashPassword(password)
	err = u.Store(email, h)
	return
}

func (u *UserUsecase) FetchJWT(email string, password string) (tokenstring string, err error) {
	id, err := u.CheckExistUser(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now(),
	})
	tokenstring, err = token.SignedString([]byte("foobar"))

	if err != nil {
		return "", err
	}
	return
}

func createHashPassword(password string) (hashPassword string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPassword = string(hash)
	return
}
