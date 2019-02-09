package mail

type UserMail struct {
	MailHandler
}

func (repo *UserMail) SendConfirmValidEmail(email string, token string) (err error) {
	// TODO: あとで実装する
	return
}
