package usecase

type UserMail interface {
	SendConfirmValidEmail(email string, token string) error
}
