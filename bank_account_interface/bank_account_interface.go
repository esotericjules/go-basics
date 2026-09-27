package main

import (
	"errors"
	"fmt"
)

type AccountManager interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount > a.Balance {
		return errors.New("Cannot withdraw more than is in the bank account")
	}
	a.Balance -= amount
	return nil
}

func main() {
	account := BankAccount{
		Owner:   "Ada",
		Balance: 500,
	}

	// var manager AccountManager = account
	var manager AccountManager = &account
	manager.Deposit(200)
	fmt.Println("manager after deposit", manager)
	manager.Withdraw(100)
	fmt.Println("manager after withdrawal", manager)

}
