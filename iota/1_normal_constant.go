package main

import "fmt"

// Normally, constants can be declared by specifying its value,
const (
	SPADES   = 1
	HEARTS   = 2
	DIAMONDS = 3
	CLUBS    = 4
)

func main() {
	fmt.Println("SPADES:", SPADES, "HEARTS:", HEARTS, "DIAMONDS:", DIAMONDS, "CLUBS:", CLUBS)
}

// $ Output:
// $ go run 1_normal_constant.go
// SPADES: 1 HEARTS: 2 DIAMONDS: 3 CLUBS: 4
