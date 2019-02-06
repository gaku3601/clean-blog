package usecase

import (
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
