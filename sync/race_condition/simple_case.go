package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
}

func main() {
	fmt.Println("Starting the program")
	var numOfRoutine int = 3
	var wg sync.WaitGroup

	wg.Add(numOfRoutine)
	for i := 1; i <= numOfRoutine; i++ {
		go func(index int) {
			fmt.Printf("Running goroutine %d\n", index)
			increment()
			wg.Done()
		}(i)
	}
	//time.Sleep(4)
	wg.Wait()
}
