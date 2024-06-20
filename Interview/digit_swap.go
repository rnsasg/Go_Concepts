package main

import (
	"fmt"
	"strconv"
)

func maximumDigitInterchange(num int) int {
	numStr := strconv.Itoa(num)
	numChars := []rune(numStr)
	n := len(numChars)
	swaps := 2

	for swaps > 0 {
		maxIdx := -1
		for i := 0; i < n-1 && swaps > 0; i++ {
			maxDigit := numChars[i]
			for j := i + 1; j < n; j++ {
				if numChars[j] > maxDigit {
					maxDigit = numChars[j]
					maxIdx = j
				}
			}
			if maxIdx != -1 && numChars[i] != numChars[maxIdx] {
				numChars[i], numChars[maxIdx] = numChars[maxIdx], numChars[i]
				swaps--
				break
			}
		}
		if maxIdx == -1 {
			break
		}
	}

	result, _ := strconv.Atoi(string(numChars))
	return result
}

func main() {
	fmt.Println(maximumDigitInterchange(2736)) // Output: 7632
	fmt.Println(maximumDigitInterchange(9876)) // Output: 9876
	fmt.Println(maximumDigitInterchange(3946)) // Output: 9643
	fmt.Println(maximumDigitInterchange(1993)) // Output: 9913
}
