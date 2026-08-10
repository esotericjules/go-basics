package main

import "fmt"

func main() {
	sum := sumup(1, 2, 3, 10)

	numbers := []int{2, 10, 15}
	anotherSum := sumup(2, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingVal int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum * startingVal

}
