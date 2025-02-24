package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Enter the revenue: ")
	if err1 != nil {
		panic(err1)
	}

	expenses, err2 := getUserInput("Enter the expenses: ")
	if err2 != nil {
		panic(err2)
	}

	taxRate, err3 := getUserInput("Enter the tax rate: ")
	if err3 != nil {
		panic(err3)
	}

	earningBeforeTax, profix, ratio := calculateFinancials(revenue, expenses, taxRate)

	// fmt.Println("ETB:", earningBeforeTax)
	fmt.Printf("ETB: %.2f\n", earningBeforeTax)
	fmt.Println("Profix:", profix)
	fmt.Println("Ratio:", fmt.Sprintf("%.2f%%", ratio))

	writeResultToFile(earningBeforeTax, profix, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := calculateEarninBeforeTax(revenue, expenses)
	profix := calculateProfix(earningBeforeTax, taxRate)
	ratio := calculateRatio(earningBeforeTax, profix)
	return earningBeforeTax, profix, ratio
}

func calculateEarninBeforeTax(revenue, expenses float64) float64 {
	return revenue - expenses
}

func calculateProfix(earningBeforeTax, taxRate float64) float64 {
	return earningBeforeTax * (1 - taxRate/100)
}

func calculateRatio(earningBeforeTax, profix float64) float64 {
	return earningBeforeTax / profix
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput < 0 {
		return 0, errors.New("invalid amount, amount should be greater than 0")
	}

	return userInput, nil
}

func writeResultToFile(earningBeforeTax, profix, ratio float64) {
	content := fmt.Sprintf("ETB: %.2f\nProfix: %.2f\nRatio: %.2f%%\n", earningBeforeTax, profix, ratio)
	os.WriteFile("result.txt", []byte(content), 0644)
}
