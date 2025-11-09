package main

import "fmt"

type floatMap map[string]float64

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println((userNames))

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9

	fmt.Println((courseRatings))

	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
