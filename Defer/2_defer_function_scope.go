package main

import "fmt"

func main() {
	defer fmt.Println("This is printed once main completes") // #4
	greet()
	fmt.Println("This is printed after greet is called") // #3
}

func greet() {
	defer fmt.Println("Printed after the first greeting") // #2
	fmt.Println("First greeting")                         // #1
}

// $ Output:
// $ go run defer_function_scope.go
// First greeting
// Printed after the first greeting
// This is printed after greet is called
// This is printed once main completes
