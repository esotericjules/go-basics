package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	fmt.Println("Student Score Tracker")
	fmt.Println("---------------------")
	passed := 0
	failed := 0
	scores := []int{}
	for _, value := range students {
		scores = append(scores, value.Score)
		result := getResult(value.Score)
		if result == "Pass" {
			passed++
		} else {
			failed++
		}
		fmt.Printf("%s scored %d: %s\n", value.Name, value.Score, result)
	}
	fmt.Println("Summary")
	fmt.Println("Passed: ", passed)
	fmt.Println("Failed: ", failed)

	averageScore(scores)

}

var students []Student = []Student{
	{
		Name:  "Ada",
		Score: 85,
	},
	{
		Name:  "John",
		Score: 48,
	},
	{
		Name:  "Mary",
		Score: 67,
	},
	{
		Name:  "Peter",
		Score: 39,
	},
}

func getResult(score int) string {
	switch {
	case score >= 50:
		return "Pass"
	default:
		return "Fail"
	}
}

func averageScore(scores []int) {
	total := 0
	for i := 0; i < len(scores); i++ {
		total = total + scores[i]
	}
	average := float64(total) / float64(len(students))
	fmt.Println("Average score: ", average)
}
