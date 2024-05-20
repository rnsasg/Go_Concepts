package main

import "fmt"

// The iota value increments within a constant declaration block only.
// So, once we declare another group of constants in a different declaration block, the iota value will reset.

const (
	SPADES = iota
	HEARTS
	DIAMONDS
	CLUBS
)

// iota resets here
const (
	ZERO = 10 * iota
	TEN
	TWENTY
	THIRTY
)

// iota resets here
const (
	BYTE     = 1 << (iota * 10) // = 1 << 0
	KILOBYTE                    // = 1 << 10
	MEGABYTE                    // = 1 << 20
	GIGABYTE                    // = 1 << 30
)

func main() {
	fmt.Println("SPADES:", SPADES, "HEARTS:", HEARTS, "DIAMONDS:", DIAMONDS, "CLUBS:", CLUBS)
	fmt.Println("ZERO:", ZERO, "TEN:", TEN, "TWENTY:", TWENTY, "THIRTY:", THIRTY)
	fmt.Println("BYTE:", BYTE, "KILOBYTE:", KILOBYTE, "MEGABYTE:", MEGABYTE, "GIGABYTE:", GIGABYTE)
}
