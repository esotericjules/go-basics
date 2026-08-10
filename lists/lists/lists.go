package lists

import "fmt"

func main() {
	prices := []float64{10.99, 9.99}
	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println("prices", prices)
}

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println("prices", prices[0:1])
// 	prices[1] = 7.99

// 	updatedPrices := append(prices, 5.99)
// 	fmt.Println("updatedPrices", updatedPrices)
// 	fmt.Println("prices", prices)
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[1] = "A carpet"
// 	fmt.Println(productNames)

// 	// slices are a window to an array in memory
// 	// when you create a slice you do not
// 	// copy that part of the array to save somewher in memory
// 	// the slice is a reference to a part of that same array in memory
// 	featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println("featuredPrices:", featuredPrices)
// 	fmt.Println("highlightedPrices:", highlightedPrices)

// 	// if we modify a slice, we also modify the array it
// 	// was created from
// 	featuredPrices[1] = 7.99
// 	fmt.Println("prices:", prices)

// 	// go creates metadata for every slice
// 	fmt.Println(len(featuredPrices)) // length gives no of slices in an array
// 	fmt.Println(cap(featuredPrices)) // capacity gives no of slices to the right an arrray can hold

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println("featuredPrices 2:", featuredPrices)
// 	fmt.Println("highlightedPrices 2:", highlightedPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }
