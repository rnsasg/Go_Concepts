package main

import "fmt"

// define our own custom constraint using an interface.
// Our interface should define a type set containing int, float, and string
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1.5, 2.5))
	fmt.Println(sum("Hare Krishan", " Hare Rama"))
}

type SumConstriant interface {
	int | float64 | string
}

func sum[T SumConstriant](a, b T) T {
	return a + b
}
