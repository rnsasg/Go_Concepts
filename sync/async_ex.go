package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- 5
	}()

	recv := <-ch
	fmt.Println(recv)
}
