package taskUser

type User struct {
	Userid         int
	FirstName      string
	LastName       string
	passportNumber string
}

func (u *User) AddPassport(passportNumber string) {
	u.passportNumber = passportNumber

}

func (u *User) GetPassport() string {
	return u.passportNumber

}
