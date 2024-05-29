package main

import "fmt"

func main() {
	// Define the indices of 'Q' positions
	qIndices := []int{1, 3, 0, 2}

	// Define the size of the board
	size := 4

	// Construct and print each string
	for _, index := range qIndices {
		format := "%" + fmt.Sprintf("%d", size) + "s\n"
		fmt.Printf(format, constructString(index, size))
	}
}

// Construct a string with 'Q' at the specified index and '.' elsewhere
func constructString(index, size int) string {
	s := make([]rune, size)
	for i := range s {
		if i == index {
			s[i] = 'Q'
		} else {
			s[i] = '.'
		}
	}
	return string(s)
}
