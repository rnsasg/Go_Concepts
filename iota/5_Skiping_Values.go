package main

import "fmt"

// iota values increment for every constant declaration.
// So, if we want to skip a value, we can declare an empty constant value:
const (
	SPADES = iota
	HEARTS
	_ // The value of iota will increase, since this is still a constant declaration
	DIAMONDS
	CLUBS
)

func main() {
	fmt.Println("SPADES:", SPADES, "HEARTS:", HEARTS, "DIAMONDS:", DIAMONDS, "CLUBS:", CLUBS)
}

// $ Output:
// $ go run 5_Skiping_Values.go
// SPADES: 0 HEARTS: 1 DIAMONDS: 3 CLUBS: 4
