package main

import (
	"fmt"
	"sync"
)

var rwmutex sync.RWMutex
var wg sync.WaitGroup

var sharedData int

func routine_1() {
	rwmutex.Lock()
	sharedData++
	rwmutex.Unlock()
	wg.Done()
}

func routine_2() {
	rwmutex.RLock()
	fmt.Println("Shared Data Value :", sharedData)
	rwmutex.RUnlock()
	wg.Done()
}

func routine_3() {
	rwmutex.Lock()
	sharedData++
	rwmutex.Unlock()
	wg.Done()
}

func main() {

	wg.Add(1)
	go routine_1()
	wg.Add(1)
	go routine_2()
	wg.Add(1)
	go routine_3()
	wg.Add(1)
	go routine_2()
	wg.Wait()
	fmt.Println(sharedData)

}
