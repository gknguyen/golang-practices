package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(len(prices), cap(prices))

	updatedPrices := append(prices, 11.99)
	fmt.Println(updatedPrices, prices)
}
