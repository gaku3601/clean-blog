package mail

type UserMail struct {
	MailHandler
}

func (repo *UserMail) ISendConfirmValidEmail(email string, token string) (err error) {
	// TODO: あとで実装する
	return
}

func (repo *UserMail) ISendForgotPasswordMail(email string, token string) (err error) {
	// TODO: あとで実装する
	return
}
