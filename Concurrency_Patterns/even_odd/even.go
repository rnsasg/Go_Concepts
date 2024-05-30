package main

import "fmt"

func generateEvenNumbers(evenChan chan<- int) {
	for i := 0; i <= 10; i += 2 {
		evenChan <- i
	}
	close(evenChan)
}

func generateOddNumbers(oddChan chan<- int) {
	for i := 1; i <= 10; i += 2 {
		oddChan <- i
	}
	close(oddChan)
}

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	go generateEvenNumbers(evenChan)
	go generateOddNumbers(oddChan)

	for evenNum := range evenChan {
		fmt.Println("Even Number:", evenNum)
	}

	for oddNum := range oddChan {
		fmt.Println("Odd Number:", oddNum)
	}
}
