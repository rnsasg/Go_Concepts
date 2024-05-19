package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Println("Index -", index)
		}(i)
	}
	wg.Wait()
}
