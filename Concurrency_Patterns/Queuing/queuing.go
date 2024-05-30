package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	limit = 2
	work  = 10
)

func main() {

	var wg sync.WaitGroup
	fmt.Println("Queue limit:", limit)
	queue := make(chan struct{}, limit)
	wg.Add(work)

	for i := 1; i <= work; i++ {
		process(i, queue, &wg)
	}

	wg.Wait()
	close(queue)

	fmt.Println("Work complete")

}

func process(work int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}
	go func() {
		defer wg.Done()

		time.Sleep(1 * time.Second)
		fmt.Println("Processed:", work)
		<-queue
	}()
}
