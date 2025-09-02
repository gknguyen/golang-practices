package main

import (
	"fmt"
	"struct/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	adminUser, err := user.NewAdmin("admin@example.com", "securepassword")
	if err != nil {
		fmt.Println("Error creating admin user:", err)
		return
	}
	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()
}

func getUserData(text string) string {
	var input string
	fmt.Print(text)
	fmt.Scanln(&input)
	return input
}
