package usecase

import "github.com/gaku3601/clean-blog/src/domain"

type UserRepository interface {
	Store(domain.User) error
}