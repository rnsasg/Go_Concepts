package main

import (
	"fmt"
	"sync"
)

func main() {
	i1 := generateWork([]int{1, 3, 5, 7, 9})
	i2 := generateWork([]int{2, 4, 6, 8, 10})

	out := fan_in(i1, i2)

	for value := range out {
		fmt.Println("Received jobs from : ", value)
	}
}

func fan_in(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(inputs))
	fmt.Println("size of channel", len(inputs))

	for idx, in := range inputs {
		go func(ch <-chan int) {
			fmt.Println("Starting a goroutine:", idx)
			for {
				value, ok := <-ch
				if !ok {
					wg.Done()
					fmt.Println("Exiting a goroutine", idx)
					break
				}
				fmt.Println("Putting value at out chhanel:", idx, value)
				out <- value
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
		fmt.Println("Closing out fan out channel")
	}()

	return out
}

func generateWork(work []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, w := range work {
			out <- w
		}
	}()

	return out
}
