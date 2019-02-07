package database

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) Store(email string, password string) (err error) {
	err = repo.InsertUser(email, password)
	return
}

func (repo *UserRepository) CheckExistUser(email string, password string) (id int, err error) {
	// TODO: 実装する
	return 1, nil
}
