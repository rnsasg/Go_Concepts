package main

import "fmt"

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	ch := generator()

	for i := 0; i < 10; i++ {
		fmt.Println("Received value from channel %d", <-ch)
	}
}
