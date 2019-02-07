package usecase

import "github.com/gaku3601/clean-blog/src/domain"

// UserUsecase ユースケースstruct
type UserUsecase struct {
	UserRepository
}

// Add ユーザを追加します。
func (u *UserUsecase) Add(email string, password string) (err error) {
	err = u.Store(email, password)
	return
}

func (u *UserUsecase) FetchJWT(email string, password string) string {
	d, _ := domain.NewUser(1, email, password)
	return d.JWT
}
