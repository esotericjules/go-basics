package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate usser input
// => show error message & exit if invalid input is provided
// - No negative numbers
// - Not O
// 2) store calculated results in a file

const profitCalculatorFile = "profit_calculator.txt"

func main() {
	revenue, err1 := getUserInput("Revenue: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	expenses, err2 := getUserInput("Expenses: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	taxRate, err3 := getUserInput("Tax Rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// data := fmt.Sprintf("%.1f\n%.1f\n%.1f\n", ebt, profit, ratio)
	// os.WriteFile(profitCalculatorFile, []byte(fmt.Sprint(data)), 0644)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResults(ebt, profit, ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	isInvalid, errorMessage := validateInput(userInput)
	if isInvalid {
		// fmt.Println(errorMessage)
		return 0, errors.New(errorMessage)
	}
	// validateInput(userInput)
	return userInput, nil
}

func validateInput(input float64) (bool, string) {
	var errorMessage string
	var isInvalid bool
	if input <= 0 {
		// errors.New("The provided input must be greater than 0 ")
		errorMessage = "The provided input must be greater than 0 "
		isInvalid = true
		// panic("The provided input must be greater than 0 ")
	}
	return isInvalid, errorMessage
}

func storeResults(ebt, profit, ratio float64) {
	data := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)

	os.WriteFile(profitCalculatorFile, []byte(data), 0644)
}
