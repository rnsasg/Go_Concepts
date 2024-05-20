package main

import "fmt"

// We can use iota and avoid explicitly declaring the incremental integer values:
const (
	SPADES   = iota
	HEARTS   = iota
	DIAMONDS = iota
	CLUBS    = iota
)

func main() {
	fmt.Println("SPADES:", SPADES, "HEARTS:", HEARTS, "DIAMONDS:", DIAMONDS, "CLUBS:", CLUBS)
}

// $ Output:
// $ go run 2_multiple_iota.go
// SPADES: 0 HEARTS: 1 DIAMONDS: 2 CLUBS: 3
