package main

import (
	"fmt"
	"sync"
)

var shared_data int
var mutex sync.Mutex
var wg sync.WaitGroup

func routine_1() {
	mutex.Lock()
	shared_data++
	mutex.Unlock()
	wg.Done()
}

func routine_2() {
	mutex.Lock()
	shared_data++
	mutex.Unlock()
	wg.Done()
}

func main() {
	wg.Add(1)
	go routine_1()
	wg.Add(1)
	go routine_2()
	wg.Wait()
	fmt.Println("Shared Data Value", shared_data)
}
