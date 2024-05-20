package main

import "fmt"

// iota starts from 0, but we can change the base value by adding the given quantity to the first declaration.
// So, if we want the same result as the first example, where SPADES = 1, then we can declare our constants as follows:
// In general, if we want to start from an index n, we can start the constant declaration with iota + <n>
const (
	// Now, the index starts from 1
	SPADES = iota + 1
	HEARTS
	DIAMONDS
	CLUBS
)

func main() {
	fmt.Println("SPADES:", SPADES, "HEARTS:", HEARTS, "DIAMONDS:", DIAMONDS, "CLUBS:", CLUBS)
}

// $ Output:
// $ go run 4_Non_Zero_Start.go
// SPADES: 1 HEARTS: 2 DIAMONDS: 3 CLUBS: 4
