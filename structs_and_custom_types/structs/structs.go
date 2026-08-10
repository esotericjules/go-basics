package main

import (
	"fmt"
	"go-basics/structs_and_custom_types/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	// assigning value to variable using instance of struct
	// curly braces in front a stuct is called a "Struct Literal" or "Composite Literal"
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	// creating appUser using the constructor function
	// appUser = newUser(userFirstName, userLastName, userBirthdate)

	// de-referencing a constructor function
	// appUser = *user.newUser(userFirstName, userLastName, userBirthdate)

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// creating new admin struct
	admin := user.NewAdmin("test@example.com", "test123")
	//accessing methods on the user struct therough admin struct
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// because outputUserDetails is now a method in user struct
	// the user value do not need to be passed to it as an argument.
	appUser.OutputUserDetails()

	appUser.ClearUserName()

	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// with scanln you can use the enter key to enter an empty input
	fmt.Scanln(&value)
	return value
}
