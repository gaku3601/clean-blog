package usecase

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/clean-blog/src/domain"
)

// UserUsecase ユースケースstruct
type UserUsecase struct {
	Repo UserRepository
}

// Add ユーザを追加します。
func (u *UserUsecase) Add(d domain.User) (err error) {
	err = u.Repo.Store(d)
	return
}

// CreateJWT JWTトークンを作成します。
func (u *UserUsecase) CreateJWT(d domain.User) (auth *domain.Auth) {
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
	return &domain.Auth{Token: tokenstring}
}
