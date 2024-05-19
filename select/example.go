package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	done := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		ch <- "Hello"
		wg.Done()
	}()

	go func() {
		ch <- "World"
		wg.Done()
	}()

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			fmt.Println("Channel Closed")
	// 			return
	// 		case msg1 := <-ch:
	// 			fmt.Println(msg1 + " ")
	// 		default:
	// 			fmt.Println("There is no message left on channel ")
	// 		}
	// 	}
	// }()

	wg.Wait()
	close(done)
	time.Sleep(5 * time.Second)
}
