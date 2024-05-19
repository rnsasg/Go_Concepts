package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32
var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go func(counter *int32) {
		atomic.AddInt32(counter, 1)
		wg.Done()
	}(&counter)

	wg.Add(1)
	go func(counter *int32) {
		atomic.AddInt32(counter, 3)
		wg.Done()
	}(&counter)

	fmt.Println("Counter Value", counter)
	wg.Wait()
}
