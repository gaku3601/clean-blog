package usecase

import "github.com/gaku3601/clean-blog/domain"

type UserUsecase struct {
	Repo UserRepository
}

func (u *UserUsecase) Add(d domain.User) (err error) {
	err = u.Repo.Store(d)
	return
}
