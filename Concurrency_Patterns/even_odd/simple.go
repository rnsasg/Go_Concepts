package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing Job %d\n", id, job)
	}
}

func main() {
	numWorker := 3
	numJobs := 10
	var wg sync.WaitGroup
	jobs := make(chan int, numJobs)

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
