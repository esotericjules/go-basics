package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1.
	var hobbies = [3]string{"Reading", "cycling", "padel"}
	fmt.Println("hobbies", hobbies)

	// 2.
	fmt.Println("first hobby", hobbies[0])
	fmt.Println("rest of hobby", hobbies[1:3])

	// 3.
	firstTwoArrSlice := hobbies[0:2]
	fmt.Println("firstTwoArrSlice", firstTwoArrSlice)

	firstTwoArrSlice2 := []string{firstTwoArrSlice[0], firstTwoArrSlice[1]}
	fmt.Println("firstTwoArrSlice2:", firstTwoArrSlice2)

	// 4.
	firstTwoArrSlice = firstTwoArrSlice[1:3]
	fmt.Println("firstTwoArrSlice again", firstTwoArrSlice)

	// 5.
	courseGoals := []string{"work on prod go app", "become fullstack"}
	fmt.Println("courseGoals", courseGoals)

	// 6.
	courseGoals[1] = "build my own apps"
	UpdatedCourseGoals := append(courseGoals, "launch my product")

	fmt.Println("UpdatedCourseGoals", UpdatedCourseGoals)

	// 7.
	products := []Product{{title: "hshd", id: 1, price: 40.7}}
	products = append(products, Product{title: "erer", id: 2, price: 50.3})
	fmt.Println("products:", products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
