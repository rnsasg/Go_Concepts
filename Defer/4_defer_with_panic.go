package main

import "fmt"

// In Go, a function panics if something went wrong. In this case, the function exits immediately,
// and doesn’t execute any statement after the panic function call.
// If any function had been deferred before panic was called, it would be executed before the function exits:

func main() {
	defer fmt.Println("This is printed once the function panics")

	fmt.Println("This is printed before panic")
	panic("Somthing Went Wrong")
	fmt.Println("This should not be printed")
}

// $ Output
// $ go run defer_with_panic.go
// This is printed before panic
// This is printed once the function panics
// panic: Somthing Went Wrong

// goroutine 1 [running]:
// main.main()
//         /Users/kumarro/go/src/Go_Concepts/Defer/defer_with_panic.go:13 +0xaa
// exit status 2
