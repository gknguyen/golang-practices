package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}

	prices := [4]float64{1, 2, 3, 4}

	fmt.Println(prices)
	fmt.Println(prices[2])

	fmt.Println(productNames)

	productNames[2] = "A Table"

	fmt.Println(productNames)

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]

	fmt.Println(highlightedPrices)

	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
