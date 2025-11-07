package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(len(prices), cap(prices))

	updatedPrices := append(prices, 11.99)
	fmt.Println(updatedPrices, prices)

	discountPrices := []float64{1.99, 2.99}

	totalPrices := append(prices, discountPrices...)
	fmt.Println(totalPrices)

}
