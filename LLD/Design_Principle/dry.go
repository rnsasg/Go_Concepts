package main

import (
	"fmt"
	"time"
)

// 1. "Every piece of knowledge must have a single, unambiguous, authoritative representation within a system."
// 2. The DRY principle encourages developers to write modular, reusable code and
// avoid duplicating the same functionality in multiple places.

// 3. It encourages us to minimize redundancy and write code that does one thing well,
// making our lives (and the lives of those who maintain our code) much easier.

// Example 1: Avoiding Code Duplication

func calculate_tax_food(price int) int {
	var tax_rate float32 = 0.07
	return price * (1 + int(tax_rate))
}

func calculate_tax_clothing(price int) int {
	var tax_rate float32 = 0.12
	return price * (1 + int(tax_rate))
}

func calculate_tax_electronics(price int) int {
	var tax_rate float32 = 0.15
	return price * (1 + int(tax_rate))
}

// Improves One

func calculate_tax(price int, tax_rate float32) int {
	return price * (1 + int(tax_rate))
}

// Why is DRY Important?
// i. Reduced Code Duplication
// ii. Improved Code Reusability
// iii. Easier Bug Fixing
// iv. Improved Consistency
// v. Faster Development

// Example 2: Using Decorators for Cross-cutting Concerns

// Function type for a function that takes two integers and returns an integer
type Operation func(int, int) int

// Logger is a higher-order function that wraps an Operation with logging
func Logger(op Operation) Operation {
	return func(a, b int) int {
		fmt.Printf("Calling function with args (%d, %d)\n", a, b)
		start := time.Now()
		result := op(a, b)
		duration := time.Since(start)
		fmt.Printf("Function returned %d, took %v\n", result, duration)
		return result
	}
}

// Some example functions that match the Operation type
func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	// Wrapping the Add and Multiply functions with the Logger
	loggedAdd := Logger(Add)
	loggedMultiply := Logger(Multiply)

	// Calling the wrapped functions
	resultAdd := loggedAdd(2, 3)
	resultMultiply := loggedMultiply(4, 5)

	fmt.Printf("Result of loggedAdd: %d\n", resultAdd)
	fmt.Printf("Result of loggedMultiply: %d\n", resultMultiply)

	// Reusability
	fmt.Println(calculate_tax(10, 0.12))
}
