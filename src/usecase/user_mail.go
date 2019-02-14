package usecase

type UserMail interface {
	SendConfirmValidEmail(email string, token string) error
	SendForgotPasswordMail(email string, token string) error
}
