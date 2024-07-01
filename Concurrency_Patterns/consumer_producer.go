package main

import (
	"fmt"
	"sync"
)

func main() {
	consumers := 2
	producerChannel := make(chan int, 20)
	consumerChannel := make(chan int, 20)

	var wg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Producer goroutine
	wg.Add(1)
	go func(producerChannel chan<- int) {
		defer wg.Done()
		for i := 1; i <= 20; i++ {
			producerChannel <- i
		}
		close(producerChannel)
	}(producerChannel)

	// Consumer goroutines
	consumerWg.Add(consumers)
	for i := 0; i < consumers; i++ {
		go func(producerChannel <-chan int, consumerChannel chan<- int) {
			defer consumerWg.Done()
			for number := range producerChannel {
				consumerChannel <- 2 * number
			}
		}(producerChannel, consumerChannel)
	}

	// Printer goroutine
	wg.Add(1)
	go func(consumerChannel <-chan int) {
		defer wg.Done()
		for number := range consumerChannel {
			fmt.Println(number)
		}
	}(consumerChannel)

	// Wait for all consumers to finish, then close consumerChannel
	go func() {
		consumerWg.Wait()
		close(consumerChannel)
	}()

	// Wait for all producer and printer goroutines to finish
	wg.Wait()
}
