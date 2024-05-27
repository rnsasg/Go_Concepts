package main

import "fmt"

func main() {
	fmt.Println(zeroOrOne())
}

// Deferred statements are executed before the functions value is returned.
// In this case, the named return value i is incremented once the function completes,
// but before the value is returned to the calling statement in the main function.

func zeroOrOne() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

// $ Output:
// $ go run defer_outout.go
// 1
