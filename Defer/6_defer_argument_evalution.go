package main

import "fmt"

var i int

func count() int {
	i++
	return i
}
func main() {
	defer func(int) {
		// count()
	}(count())
	fmt.Println("current count:", count())
}

// func printCount() {
// 	fmt.Println("count:", count())
// }

// $ Output:
// $ go run defer_argument_evalution.go
// current count: 1
// count: 2
