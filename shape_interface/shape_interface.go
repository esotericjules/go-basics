package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func PrintArea(s Shape) {
	fmt.Printf("%T: %.2f\n", s, s.Area())
}

func main() {
	rectangle := Rectangle{Height: 20, Width: 20}
	circle := Circle{Radius: 5}
	PrintArea(rectangle)
	PrintArea(circle)

	shapes := []Shape{
		Rectangle{Height: 20, Width: 20},
		Circle{Radius: 5},
		Rectangle{Height: 4, Width: 4},
		Triangle{Base: 10, Height: 5},
	}

	fmt.Println(shapes)
	var area float64
	for index, shape := range shapes {
		fmt.Println(index, shape)
		area += shape.Area()
	}
	fmt.Println("combined area =", area)
}
