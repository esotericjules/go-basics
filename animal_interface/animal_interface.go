package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Sound() string
	Move() string
}

type Dog struct {
	Name string
}

type Bird struct {
	Name string
}

type Fish struct {
	Name string
}

func (d Dog) Sound() string {
	return "Woof"
}

func (d Dog) Move() string {
	return "Running"
}

func (b Bird) Sound() string {
	return "Chirp"
}

func (b Bird) Move() string {
	return "flying"
}
func (f Fish) Sound() string {
	return "swish"
}

func (f Fish) Move() string {
	return "swimming"
}

func DescribeAnimal(animal Animal) {
	animalType := fmt.Sprintf("%T", animal)
	fmt.Println(strings.TrimPrefix(animalType, "main."))
	fmt.Println("Sound:", animal.Sound())
	fmt.Println("Movement:", animal.Move())
}

func main() {

	dog := Dog{Name: "Bingo"}
	fish := Fish{Name: "Fish"}
	DescribeAnimal(dog)
	DescribeAnimal(fish)

}
