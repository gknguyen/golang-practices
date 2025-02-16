package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Wht do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	for {
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount < 0 {
				fmt.Println("Invalid amount, deposit amount should be greater than 0!")
			} else {
				accountBalance += depositAmount
				fmt.Println("Updated! Your new balance is:", accountBalance)
			}
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount < 0 {
				fmt.Println("Invalid amount, withdraw amount should be greater than 0!")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance!")
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Updated! Your new balance is:", accountBalance)
			}
		default:
			fmt.Println("Thank you for using Go Bank!")
			return
		}
	}
}
