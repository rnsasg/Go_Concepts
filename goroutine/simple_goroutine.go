package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Go Routine!!")
}

func main() {
	// Create a new goroutine that executes the sayHello function.
	go sayHello()
	// Allow some time for the goroutine to execute.
	time.Sleep(1)
	fmt.Println("Hello from main program !!")
}
