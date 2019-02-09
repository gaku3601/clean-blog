package mail

import "github.com/gaku3601/clean-blog/src/interfaces/mail"

type MailHandler struct {
}

func NewMailHandler() mail.MailHandler {
	return new(MailHandler)
}
