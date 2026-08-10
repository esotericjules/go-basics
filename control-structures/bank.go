package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

//used to write data to a file

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to GO bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("---------")
			// panic("Cant continue, sorry")
		}

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount must be greater than 0 ")
				continue
				// return
			}

			accountBalance += depositAmount
			fmt.Println("Your balance is now: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("Amount to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount must be less or equal to accountBalance")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your balance is now: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Unfortunately nothing can bde done for you")
			fmt.Println("Thanks for choosing our bank")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount must be greater than 0 ")
		// 		continue
		// 		// return
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Your balance is now: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Println("Amount to withdraw: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)

		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount must be less or equal to accountBalance")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Your balance is now: ", accountBalance)
		// } else {
		// 	fmt.Println("Unfortunately nothing can bde done for you")
		// 	break
		// }
	}
	// fmt.Println("Thanks for choosing our bank")
}
