package main

import "fmt"

// Although the value of iota increases by one incrementally,
// we can make use of basic math operations to modify the final value assigned to each constant.
// For example, if we use iota * 10, it will give the n’th multiple of ten for each const declaration:
const (
	ZERO   = 10 * iota // = 10 * 0
	TEN                // = 10 * 1
	TWENTY             // = 10 * 2
	THIRTY             // = 10 * 3
)

func main() {
	fmt.Println("ZERO:", ZERO, "TEN:", TEN, "TWENTY:", TWENTY, "THIRTY:", THIRTY)
}

// $ Output
// $ go run 6_Custom_Math_Op.go
// ZERO: 0 TEN: 10 TWENTY: 20 THIRTY: 30
