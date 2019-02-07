package database

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) Store(email string, password string) (err error) {
	err = repo.InsertUser(email, password)
	return
}

func (repo *UserRepository) CheckExistUser(email string, password string) (id int, err error) {
	id, err = repo.FetchUserID(email, password)
	return
}
