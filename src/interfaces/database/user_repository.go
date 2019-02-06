package database

import "github.com/gaku3601/clean-blog/src/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (err error) {
	err = repo.InsertUser(u.Email, u.Password)
	return
}
