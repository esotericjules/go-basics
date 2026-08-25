package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Person struct {
	Name string
}

type Robot struct {
	id int
}

func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func (d Dog) Speak() string {
	return d.Name + " says: Woof"
}

func (p Person) Speak() string {
	return p.Name + " says: Hello there!"
}

func main() {
	dog := Dog{Name: "Bingo"}
	person := Person{Name: "Juju"}
	MakeSpeak(dog)
	MakeSpeak(person)
}
