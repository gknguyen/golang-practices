package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	moreNumbers := []int{5, 1, 2}

	transformedNumbers := transformNumbers(&numbers, getTransformerFunction(&numbers))
	fmt.Println(transformedNumbers)
	moreTransformedNumbers := transformNumbers(&moreNumbers, getTransformerFunction(&moreNumbers))
	fmt.Println(moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
