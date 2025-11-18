package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15, 40, -5}
	sum := sumup(numbers...)

	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
