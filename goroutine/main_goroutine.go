package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Hello from go routine !!!")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Hello from main go routine !!")
}
