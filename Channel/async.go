package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)

	go func() {
		fmt.Println("Waiting")
		//time.Sleep(10 * time.Second)
		ch <- true
		ch <- false
	}()
	time.Sleep(2 * time.Second)
	select {
	case msg := <-ch:
		fmt.Println("Message received ", msg)
	default:
		fmt.Printf("Closing the channel")
	}
	fmt.Println("Exiting main")
}
