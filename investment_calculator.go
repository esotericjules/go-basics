package main

import (
	"fmt"
	"math"
)

// with constants you have to assign an initial value
// because they cannot be changed
const inflationRate = 2.5

// the main function contains the code that will be
// executed when the application starts
func main1() {

	var investmentAmount float64 = 1000

	// shorthand syntax. widely used. type is inferred
	expectedReturnRate := 5.5

	// variable declaration syntax without assignment
	// here the type must be declared so GO knows it
	var years float64

	fmt.Print("Investment Amount, years, expectedReturnRate : ")

	// this function scans the terminal for input
	// to use a variable as an argument to the scan function you
	// have to use a pointer to reference the variable.
	// & serves as a pointer. This allows scan to populate the variable as a value
	fmt.Scan(&investmentAmount, &years, &expectedReturnRate)
	// fmt.Scan(&years)
	// fmt.Scan(&expectedReturnRate)

	// var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := futureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Real Future Value: %.1f\n", futureRealValue)

	fmt.Printf("%T\n", investmentAmount)

	// with Println we can enter a new line after every value
	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

}

func main() {
	profitCalculator()
}

func futureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

func profitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue, Expenses, Tax Rate: ")
	fmt.Scan(&revenue, &expenses, &taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax: ", profit)
	fmt.Println("ratio: ", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses

	taxAmount := earningsBeforeTax * taxRate / 100

	earningsAfterTax := earningsBeforeTax - taxAmount
	ratio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
