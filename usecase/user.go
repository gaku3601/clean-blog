package usecase

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/clean-blog/domain"
)

type UserUsecase struct {
	Repo UserRepository
}

func (u *UserUsecase) Add(d domain.User) (err error) {
	err = u.Repo.Store(d)
	return
}

func (u *UserUsecase) CreateJWT(d domain.User) string {
	// User情報をtokenに込める
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"ID":    d.ID,
		"Email": d.Email,
	})
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
