package domain

type User struct {
	ID            int
	Email         string
	Password      string
	ValidEmail    bool
	ValidPassword bool
}
