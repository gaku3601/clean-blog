package usecase

type UserInterface interface {
	AddUser(email string, password string) (err error)
	ReSendConfirmValidEmail(email string) (err error)
	ChangeUserPassword(id int, password string, nextPassword string) (err error)
	GetAccessToken(email string, password string) (token string, err error)
	ActivationEmail(validToken string) (err error)
	ActivationPassword(id int, password string) (err error)
	ForgotPassword(email string) (err error)
	ProcessForgotPassword(token string, newPassword string) (err error)
	CertificationSocialProfile(service ServiceEnum, email string, uid string) (token string, err error)
}
