package main

import "fmt"

func sum[T any](a, b T) T {
	return (a + b)
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(4.0, 2.0))
	fmt.Println(sum("a", "b"))
}

// # While constraint of type any generally works it does not support operators.
// $ go run main.go
// ./any_type.go:6:10: invalid operation: operator + not defined on a (variable of type T constrained by any)
