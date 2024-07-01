package main

import "fmt"

type constraints interface {
	int | float64 | string
}

func Add[T constraints](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.1, 2.2))

}

// package main

// import "fmt"

// func main() {
// 	willPanic()
// }

// func handlePanic() {
// 	data := recover()
// 	fmt.Println("Recovered data", data)
// }

// func willPanic() {
// 	defer handlePanic()
// 	panic("started")
// }

// package main

// import "fmt"

// func generator(work []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for w := range work {
// 			out <- w
// 		}
// 	}()
// 	return out
// }

// func fan_out(input <-chan int) <-chan int {

// 	out := make(chan int)

// 	go func() {
// 		defer close(out)
// 		for data := range input {
// 			out <- data
// 		}
// 	}()

// 	return out
// }

// func main() {
// 	gen := generator([]int{1, 2, 3, 4, 5, 6, 7, 8})

// 	out1 := fan_out(gen)
// 	out2 := fan_out(gen)
// 	out3 := fan_out(gen)
// 	out4 := fan_out(gen)

// 	for {
// 		select {
// 		case val1 := <-out1:
// 			fmt.Printf("Received value %d from channel-1\n", val1)
// 		case val2 := <-out2:
// 			fmt.Printf("Received value %d from channel-2\n", val2)
// 		case val3 := <-out3:
// 			fmt.Printf("Received value %d from channel-3\n", val3)
// 		case val4 := <-out4:
// 			fmt.Printf("Received value %d from channel-4\n", val4)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func generator(work []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for w := range work {
// 			out <- w
// 		}
// 	}()
// 	return out
// }

// func fan_in(inputs ...<-chan int) <-chan int {
// 	var wg sync.WaitGroup
// 	out := make(chan int)
// 	wg.Add(len(inputs))

// 	for idx, in := range inputs {
// 		go func(ch <-chan int) {
// 			fmt.Println("Starting a goroutine:", idx)
// 			for {
// 				val, ok := <-ch
// 				if !ok {
// 					wg.Done()
// 					fmt.Printf("Closing go routine : %d\n", idx)
// 					break
// 				}
// 				fmt.Println("Putting value at out chhanel:", idx, val)
// 				out <- val
// 			}
// 		}(in)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

// func main() {
// 	g1 := generator([]int{1, 2, 3, 4, 5, 6})
// 	g2 := generator([]int{1, 2, 3, 4, 5, 6})

// 	out := fan_in(g1, g2)

// 	for val := range out {
// 		fmt.Println(val)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	work := []int{0, 1, 2, 3, 4, 5, 6}
// 	gen := generator(work)
// 	fil := filter(gen)
// 	sq := square(fil)
// 	half := square(sq)

// 	for f := range half {
// 		fmt.Println(f)
// 	}
// }

// func filter(ch <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for r := range ch {
// 			if r%2 == 0 {
// 				out <- r
// 			}
// 		}
// 	}()
// 	return out
// }

// func square(ch <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for r := range ch {
// 			out <- r * r
// 		}
// 	}()
// 	return out
// }

// func half(ch <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for r := range ch {
// 			out <- r / 2
// 		}
// 	}()
// 	return out
// }

// func generator(work []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for w := range work {
// 			out <- w
// 		}
// 	}()
// 	return out
// }

// package main
// import "fmt"

// func generator(work []int) <-chan int {
// 	out := make(chan int)

// 	go func() {
// 		defer close(out)
// 		for w := range work {
// 			out <- w
// 		}
// 	}()

// 	return out
// }

// func main() {
// 	work := []int{1, 2, 3, 4, 5}
// 	ch := generator(work)
// 	for range work {
// 		fmt.Println(<-ch)
// 	}
// }

// package main

// import "fmt"

// func generator() <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; ; i++ {
// 			ch <- i
// 		}
// 	}()
// 	return ch
// }

// func main() {
// 	ch := generator()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(<-ch)
// 	}
// }
