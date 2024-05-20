package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	panic("Something went wrong")
	fmt.Println("This should not be printed")
}

// $ Output:
// $ go run defer_panic_recover.go
// Recovered from panic Something went wrong
