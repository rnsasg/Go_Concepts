package main

import "math"

// https://newsletter.ashishps.com/p/21b57678-b351-4ed4-b390-3b6308af2f7d

// What is the KISS Principle?
// The KISS principle, which stands for "Keep it Simple, Stupid" is a design philosophy that
// emphasizes the importance of simplicity in software development.

// It suggests that software should be designed to be easy to understand, modify, and extend,
// rather than complex and convoluted.

// Example: Calculating Factorial
// Overly Complex:

// KISS-aligned:

//Overly Complex:

func Factorial(n int) int {
	if n < 0 {
		return -1 // Factorial is not defined for negative numbers
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// use function if available

func Factorial(n int) int {
	if n < 0 {
		return -1 // Factorial is not defined for negative numbers
	}
	return math.Factorial(n)
}
