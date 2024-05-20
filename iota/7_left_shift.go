package main

import "fmt"

// we can use the shift-left operation to make exponential increments:
const (
	BYTE     = 1 << (iota * 10) // = 1 << 0
	KILOBYTE                    // = 1 << 10
	MEGABYTE                    // = 1 << 20
	GIGABYTE                    // = 1 << 30
)

func main() {
	fmt.Println("BYTE:", BYTE, "KILOBYTE:", KILOBYTE, "MEGABYTE:", MEGABYTE, "GIGABYTE:", GIGABYTE)
}

// $ Output:
// $ go run 7_left_shift.go
// BYTE: 1 KILOBYTE: 1024 MEGABYTE: 1048576 GIGABYTE: 1073741824
