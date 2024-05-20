package main

import (
	"fmt"
	"sync"
)

func generateEvenNumber(evenChan chan<- int, wg *sync.WaitGroup) {
	//defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			evenChan <- i
		}
	}
	close(evenChan)
}

func generateOddNumber(oddChan chan<- int, wg *sync.WaitGroup) {
	//defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			oddChan <- i
		}
	}
	close(oddChan)
}

func main() {
	var wg *sync.WaitGroup
	evenChan := make(chan int)
	oddChan := make(chan int)

	//wg.Add(2)
	go generateEvenNumber(evenChan, wg)

	go generateOddNumber(oddChan, wg)

	for num := range evenChan {
		fmt.Println(num)
	}

	for num := range oddChan {
		fmt.Println(num)
	}
	//wg.Wait()
}
