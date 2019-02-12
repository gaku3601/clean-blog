package usecase

type UserMail interface {
	ISendConfirmValidEmail(email string, token string) error
	ISendForgotPasswordMail(email string, token string) error
}
