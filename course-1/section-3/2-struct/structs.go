package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (user *User) outputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthDate, user.createdAt)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	appUser.outputUserDetails()
}

func getUserData(text string) string {
	var input string
	fmt.Print(text)
	fmt.Scan(&input)
	return input
}
