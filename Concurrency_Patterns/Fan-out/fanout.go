package main

import "fmt"

func main() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	gen := generateWork(work)

	out1 := fanout(gen)
	out2 := fanout(gen)
	out3 := fanout(gen)
	out4 := fanout(gen)

	for {
		select {
		case val1 := <-out1:
			fmt.Println("Output 1 go", val1)
		case val2 := <-out2:
			fmt.Println("Output 2 go", val2)
		case val3 := <-out3:
			fmt.Println("Output 3 go", val3)
		case val4 := <-out4:
			fmt.Println("Output 4 go", val4)
		}
	}

}

func fanout(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for data := range input {
			out <- data
		}
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
