package annonymous

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	tripple := createTransformer(3)

	// transformed := transformNumbers(&numbers, func(number int) int {
	// 	return number * 2
	// })

	doubled := transformNumbers(&numbers, double)
	trippled := transformNumbers(&numbers, tripple)

	fmt.Println("doubled:", doubled)
	fmt.Println("trippled:", trippled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
