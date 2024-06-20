package main

import "fmt"

func createCounter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
}

func main() {
	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter2())
}
