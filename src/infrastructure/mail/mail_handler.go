package mail

import "github.com/gaku3601/clean-blog/src/usecase"

type MailHandler struct {
}

func NewMailHandler() usecase.UserMail {
	return new(MailHandler)
}

func (m *MailHandler) SendConfirmValidEmail(email string, token string) (err error) {
	return
}
func (m *MailHandler) SendForgotPasswordMail(email string, token string) (err error) {
	return
}
