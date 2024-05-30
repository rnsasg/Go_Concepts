package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func sum[T constraints.Ordered](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1.5, 2.5))
	fmt.Println(sum("Hare Krishan", " Hare Rama"))
}

// We can also use the constraints package which defines a set of useful constraints to be used with type parameters.
// type Signed interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64
// }

// type Unsigned interface {
// 	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
// }

// type Integer interface {
// 	Signed | Unsigned
// }

// type Float interface {
// 	~float32 | ~float64
// }

// type Complex interface {
// 	~complex64 | ~complex128
// }

// type Ordered interface {
// 	Integer | Float | ~string
// }
