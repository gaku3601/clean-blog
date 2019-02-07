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

func (u *UserUsecase) FetchJWT(email string, password string) (token string, err error) {
	id, err := u.CheckExistUser(email, password)
	if err != nil {
		return "", err
	}
	d, err := domain.NewUser(id, email, password)
	if err != nil {
		return "", err
	}
	token = d.JWT
	return
}
