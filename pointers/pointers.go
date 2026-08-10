package main

import "fmt"

// A pointers nill value is nil
// nil represents the absence of an address value
// That is a pointer pointing at no address/no value in memory

func main() {
	age := 32

	// & ampersand is used in front of a notmal variable
	// to get a pointer to the address of the variable
	var agePointer *int = &age

	// * asterisks is used to de-reference the address
	// converting the address to the stored value
	fmt.Println("Age", *agePointer)

	// age is passed as a pointer &age
	// because the function is expecting a pointer
	// adultYears := getAdultYears(agePointer)
	// fmt.Println(adultYears)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	// To perform an arithemic on a pointer
	// you have to de-reference it first
	// return *age - 18

	*age = *age - 18
}
