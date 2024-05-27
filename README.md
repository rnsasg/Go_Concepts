# Go_Concepts
Go Core Concepts

## What is goroutine 

```txt
A goroutine in GO is a lightweight execution unit in a go program. Similar to thread. 
```

## What are goroutines and how are they different from threads? 

1. Memory Footprint : Light Weight Vs Heavy Weight  ( < 2KB )
2. Multiplexity  : Managed by Go Runtime [ Thousands of goroutines ] Vs OS Kernel
3. Synchronization :
	mutex , semaphor --> Race conditions --> C,C++,JAVA
	Mutex, Channels , WaitGroup --> Go 

4. Erorr Handling : 
	Go --> Return Values , Channels , Recover , Defer 
	C/C++/Java --> Exception , retrun value
## Concurrency Vs Parallalism


## How do you handle panics and recover from them in Go?

Panics are unexpected errors that can occur during program execution.
You can use the recover() function inside a deferred function to catch and handle panics.


```
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	panic("Something went wrong")
	fmt.Println("This should not be printed")
```
> [Note!]
> In Go, you can handle panics using the recover function within a deferred function. When a panic occurs, the normal flow of the program is interrupted, and the runtime begins to unwind the stack, running all deferred functions (in the reverse order they were deferred) until recover is called or the program exits.

## What is the difference between defer, go, and goroutine?

defer is used to schedule a function call to be executed when the surrounding function returns. go is used to start a new goroutine for concurrent execution. goroutine is the lightweight thread managed by the Go runtime.

## Function Closure in go [link](Docs/function_closure.md)

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

## What are the differences between value receivers and pointer receivers in Go methods?

## How do you implement interfaces in Go?

Interfaces in Go are implemented implicitly. If a type satisfies all the methods of an interface, it is said to implement that interface.

There is no need to explicitly declare that a type implements an interface.

## What is the purpose of the context package in Go?

The context package provides a way to pass cancellation signals, deadlines, and other request-scoped values across API boundaries.

It is used to manage the lifecycle of operations and facilitate cancellation and timeout handling.

## What is the purpose of the init() function in Go?

The init() function is a special function in Go that is automatically called before the program starts.
It is often used for initialization tasks such as registering drivers, initializing global variables, or setting up configuration.

## How do you handle time-related operations in Go?

Go provides the time package for working with dates, times, durations, and timers.
It offers functions and methods for parsing, formatting, and manipulating time values.

## What is the difference between shallow copy and deep copy in Go?


dest := make([]int,l)
src ([]int)

copy(dest,src) --> Shallow

obj1
	obj2 
		obj3
	obj4 


obj1 --> Objnew
	

