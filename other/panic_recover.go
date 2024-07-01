package main

import "fmt"

func main() {
	willPanic()
}

func handlePanic() {
	data := recover()
	fmt.Println("Recovered data", data)
}

func willPanic() {
	defer handlePanic()
	panic("started")
}
