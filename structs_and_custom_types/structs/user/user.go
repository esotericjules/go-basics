// This is the user package with its struct and all its methods
package user

import (
	"errors"
	"fmt"
	"time"
)

// struct allows for creating custom types.
// inside them are fields
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	// embedding the user struct into the admin struct
	// this is annonymous embedding
	User
	//this is not annonymous embedding
	// User: User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

// New validates the input and returns a new User.
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// outputUserDetails function becomes a method
// inside struct and the arguments inside the first () is called a receiver
// that receives the values from the struct
func (u User) OutputUserDetails() {
	// pointers to structs has an exception that allows
	// the values to be used directly without dereferencing them
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// When you define methods that should edit a struct, the value you receive
// should be a pointer that points directly to that struct. If the value is
// not a pointer, what the receiver receives is a copy of the struct, hence edits
// to the struct will be made on just the copy so this code is wrong.
// func (u user) clearUserName() {
// 	u.firstName = " "
// 	u.lastName = " "
// }

// correct code where the receiver is a pointer to the struct
func (u *User) ClearUserName() {
	u.firstName = " "
	u.lastName = " "
}

// constructor functions are utility functions that takes care of creating struct.
// They are more like patterns in go dev world rather than a feature baked into
// the language. the convention is to start constructor functions with the word "new"
// func newUser(firstName, lastName, birthDate string) user {
// 	// this returns a copy to the user struct
// 	return user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthDate: birthDate,
// 		createdAt: time.Now(),
// 	}
// }
