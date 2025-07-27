package taskUser

import (
	"fmt"
)

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

func (u User) Insert() int {

	fmt.Println("Добавлен пользователь ", u.Userid)

	return u.Userid

}
