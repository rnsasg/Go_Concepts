package main

import "fmt"

// https://www.sohamkamani.com/golang/defer/
// The defer keyword instructs a function to execute after the surrounding function completes.
func main() {
	defer fmt.Println("This is printed once the function completes")

	fmt.Println("This is printed first")
	fmt.Println("This is printed second")
}

// $ Output:
// $ go run defer_simple.go
// This is printed first
// This is printed second
// This is printed once the function completes
