package user

import (
	"fmt"
)

type User struct {
	Userid         int
	FirstName      string
	LastName       string
	passportNumber string
}

type CreateUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
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
