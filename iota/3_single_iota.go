package main

import "fmt"

// We can also omit iota declarations after the first one, so this code would give the same result:
const (
	SPADES = iota
	HEARTS
	DIAMONDS
	CLUBS
)

func main() {
	fmt.Println("SPADES:", SPADES, "HEARTS:", HEARTS, "DIAMONDS:", DIAMONDS, "CLUBS:", CLUBS)
}

// $ Output:
// $ go run 3_single_iota.go
// SPADES: 0 HEARTS: 1 DIAMONDS: 2 CLUBS: 3
