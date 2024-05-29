package main

import (
	"fmt"
	"sync"
)

func main() {
	numberChannel := make(chan int)
	evenTurn := make(chan struct{}, 1)
	oddTurn := make(chan struct{}, 1)
	var wg sync.WaitGroup

	// Producer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 20; i++ {
			numberChannel <- i
		}
		close(numberChannel)
	}()

	// Consumer goroutines
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			case num, ok := <-numberChannel:
				if !ok {
					return
				}
				if num%2 == 0 {
					<-evenTurn
					fmt.Printf("Even Consumer: %d\n", num)
					oddTurn <- struct{}{}
				}
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case num, ok := <-numberChannel:
				if !ok {
					return
				}
				if num%2 != 0 {
					<-oddTurn
					fmt.Printf("Odd Consumer: %d\n", num)
					evenTurn <- struct{}{}
				}
			}
		}
	}()

	// Start with odd turn
	oddTurn <- struct{}{}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All tasks completed")
}
