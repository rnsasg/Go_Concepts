package main

import (
	"fmt"
	"sync"
)

const (
	totalJobs   = 4
	totalWorker = 2
)

func main() {

	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	for w := 1; w <= totalWorker; w++ {
		go worker(w, jobs, results)
	}

	//Send Jobs
	for i := 1; i <= totalJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Receive results
	for a := 1; a <= totalJobs; a++ {
		fmt.Println("Recevied Results", <-results)
	}
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			fmt.Printf("Worker id %d has started job %d\n", id, j)

			result := job * 2

			results <- result

			fmt.Printf("Worker id %d has finished the job %d\n", id, j)
		}(j)
	}

	wg.Wait()
}
