package main

import "fmt"

func main() {
	revenue := getUserInput("Enter the revenue: ")
	expenses := getUserInput("Enter the expenses: ")
	taxRate := getUserInput("Enter the tax rate: ")

	earningBeforeTax, profix, ratio := calculateFinancials(revenue, expenses, taxRate)

	// fmt.Println("ETB:", earningBeforeTax)
	fmt.Printf("ETB: %.2f\n", earningBeforeTax)
	fmt.Println("Profix:", profix)
	fmt.Println("Ratio:", fmt.Sprintf("%.2f%%", ratio))
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

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
