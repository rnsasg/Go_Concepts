package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var mutex sync.Mutex
var wg sync.WaitGroup

func increment() {

	for i := 0; i < 10; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
	wg.Done()
}

func main() {

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println(counter)
}
