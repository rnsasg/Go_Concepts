package main

import "fmt"

func main() {
	// Explicit Type mention
	// sum[int](1, 2)
	// sum[float64](1.2, 1.3)
	// sum[string]("a", "b")

	// Go 1.18 comes with type inference which helps us to write code that calls generic functions without explicit types.
	sum(1, 2)
	sum(1.2, 1.3)
	sum("a", "b")
}

func sum[T interface{}](a, b T) {
	fmt.Println(a, b)
}
