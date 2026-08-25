package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Owner        string
	Balance      float64
	Transactions int
}

type Person struct {
	Name string
}

func main() {
	Juliet := Account{
		Owner:        "Juliet",
		Balance:      1000,
		Transactions: 0,
	}

	Ada := Account{
		Owner:        "Ada",
		Balance:      500,
		Transactions: 0,
	}

	fmt.Println("Account Details")
	fmt.Print("----------------\n\n")
	printBalance(Juliet)

	fmt.Print("Depositing 250...\n\n")
	deposit(250, &Juliet)
	printBalance(Juliet)

	fmt.Print("Withdrawing 100...\n\n")
	err := withdraw(100, &Juliet)
	if err != nil {
		fmt.Printf("%s\n\n", err)
	}
	printBalance(Juliet)

	fmt.Print("Withdrawing 1500...\n\n")
	err = withdraw(1500, &Juliet)
	if err != nil {
		fmt.Printf("%s\n\n", err)
	}
	printBalance(Juliet)

	/// Optional challenge 1
	err = Juliet.Transfer(200, &Ada)
	if err != nil {
		fmt.Println(err)
	}

	var person = &Person{Name: "Juliet"}
	fmt.Println(*person)

}

func deposit(amount float64, account *Account) {
	account.Balance += amount
	account.Transactions++
	receipt(account, amount, "Deposit")
}

func withdraw(amount float64, account *Account) error {

	if amount > account.Balance {
		return errors.New("Insufficient funds")
	}

	account.Balance -= amount
	account.Transactions++

	receipt(account, amount, "Withdraw")
	return nil

}

func printBalance(account Account) {
	fmt.Printf("Owner: %s\n", account.Owner)
	fmt.Printf("Balance: %.2f\n\n", account.Balance)

}

func (source *Account) Transfer(amount float64, destinationAccount *Account) error {
	if amount > source.Balance {
		fmt.Print("Insufficient funds\n\n")
		return errors.New("Insufficient funds")
	}
	source.Balance -= amount
	source.Transactions++

	destinationAccount.Balance += amount
	destinationAccount.Transactions++

	fmt.Printf("%s transfers %.f to %s\n\n", source.Owner, amount, destinationAccount.Owner)
	receipt(source, amount, "Withdraw")
	receipt(destinationAccount, amount, "Deposit")
	return nil
}

func receipt(account *Account, amount float64, action string) {
	fmt.Println("----------------")
	fmt.Println(action)
	fmt.Printf("Amount: %.f\n", amount)
	fmt.Printf("Account Owner: %s\n", account.Owner)
	fmt.Printf("New Balance: %.f\n", account.Balance)
	fmt.Printf("Transactions: %d\n", account.Transactions)
	fmt.Print("----------------\n\n")

}
