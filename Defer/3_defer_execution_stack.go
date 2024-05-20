package main

import "fmt"

func main() {
	myFunction()
}

// When there are multiple deferred statements in the same function,
// they are stored and executed as a stack.

func myFunction() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")

	fmt.Println("Starting myFunction...")
}

// $ Output:
// $ go run defer_execution_stack.go
// Starting myFunction...
// third
// second
// first
