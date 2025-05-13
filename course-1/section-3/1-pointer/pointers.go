package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age

	fmt.Println("Age:", age)

	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Age Pointer value:", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult Age:", age)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
