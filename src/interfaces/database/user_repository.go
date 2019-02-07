package database

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(email string, password string) (err error) {
	err = repo.InsertUser(email, password)
	return
}
