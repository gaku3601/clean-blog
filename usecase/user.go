package usecase

import (
	"log"
	"time"

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
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    d.ID,
		"email": d.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now(),
	})
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
