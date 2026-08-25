package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	Name     string
	Age      int
	Position string
	Manager  *Employee
	Mentor   *Employee
}

var manager *Employee

func main() {

	Alice := Employee{
		Name:     "Alice",
		Age:      45,
		Position: "Engineering Manager",
	}

	Bob := Employee{
		Name:     "Bob",
		Age:      28,
		Position: "Backend Engineer",
	}

	Charlie := Employee{
		Name:     "Charlie",
		Age:      26,
		Position: "Frontend Engineer",
	}
	PrintEmployee(Bob, "Initial state")
	PrintEmployee(Charlie, "Initial state")

	err := Bob.AssignManager(&Alice)
	if err != nil {
		fmt.Println("------Error while assigning manager------")
		fmt.Println(err)
	} else {
		PrintEmployee(Bob, "After Assigning Manager")
	}

	err = Charlie.AssignManager(&Charlie)
	if err != nil {
		fmt.Println("------Error while assigning manager------")
		fmt.Println(err)
	} else {
		PrintEmployee(Charlie, "After Assigning Manager")

	}

	RemoveManager(&Bob)
	PrintEmployee(Bob, "After Removing Manager")

	fmt.Println(manager.Name)

}

func (employee *Employee) AssignManager(manager *Employee) error {
	if employee == manager {
		return errors.New("An employee cannot be their own manager.")
	}
	employee.Manager = manager
	return nil
}

func PrintEmployee(employee Employee, action string) {
	fmt.Printf("------%s-------\n", action)
	fmt.Printf("Name: %s\n", employee.Name)
	fmt.Printf("Position: %s\n\n", employee.Position)
	if employee.Manager == nil {
		fmt.Printf("Manager: None\n\n")
	} else {
		fmt.Printf("Manager: %s\n\n", employee.Manager.Name)
	}

}

func RemoveManager(employee *Employee) {
	employee.Manager = nil

}
