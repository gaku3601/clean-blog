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

func (repo *UserRepository) UpdateValidEmail(email string) (err error) {
	// TODO: あとで実装する
	return nil
}

func (repo *UserRepository) CreateSocialProfile(servise string, email string, uid string) error {
	// TODO: あとで実装する
	return nil
}

func (repo *UserRepository) CheckExistSocialProfile(servise string, uid string) (userID int, err error) {
	// TODO: あとで実装する
	return
}
