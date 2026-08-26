package main

import "fmt"

type Incrementer interface {
	Increment()
}

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value += 1
}
func main() {
	counter := &Counter{Value: 1}
	fmt.Println("counter:", counter)
	counter.Increment()
	fmt.Println("counter:", counter.Value)
	var i Incrementer = counter
	i.Increment()
	fmt.Println(i)

}
