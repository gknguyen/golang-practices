package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (user *User) OutputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthDate, user.createdAt)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("invalid user data")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("invalid admin data")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "admin",
			lastName:  "admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}, nil
}
