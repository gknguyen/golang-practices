package main

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	var hobbies = [3]string{"games", "music", "sport"}
	fmt.Println("1.", hobbies)

	var firstHobby = hobbies[0]
	fmt.Println("2.1", firstHobby)
	var restHobbies = hobbies[1:3]
	fmt.Println("2.2", restHobbies)

	var slide1Hobbies = hobbies[0:2]
	fmt.Println("3.1", slide1Hobbies)
	var slide2Hobbies = hobbies[:2]
	fmt.Println("3.2", slide2Hobbies)

	var reslide1Hobbies = slide1Hobbies[1:3]
	fmt.Println("4.1", reslide1Hobbies)
	var reslide2Hobbies = slide2Hobbies[1:3]
	fmt.Println("4.2", reslide2Hobbies)

	var goals = []string{"cert", "knowledge"}
	fmt.Println("5.", goals)

	goals[1] = "money"
	fmt.Println("6.1", goals)

	newGoals := append(goals, "job")
	fmt.Println("6.2", newGoals)

	var products = []Product{
		{id: "1", title: "egg", price: 100},
		{id: "2", title: "milk", price: 200},
	}
	fmt.Println("7.1", products)

	newProducts := append(products, Product{id: "3", title: "bread", price: 300})
	fmt.Println("7.2", newProducts)
}
