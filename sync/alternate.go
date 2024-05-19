package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var flip bool = true
	var i int = 1

	wg.Add(2)

	go func() {
		fmt.Println("Starting Odd Go Routine")
		for {

			if flip == true && i%2 == 1 {
				fmt.Println("Odd :", i)
				i++
				flip = false
			}

			if i == 10 {
				break
			}
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("Starting Even Go Routine")
		for {
			if flip == false && i%2 == 0 {
				fmt.Println("Even :", i)
				i++
				flip = true
			}
			if i == 11 {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Exiting Main Go Routine")
}
