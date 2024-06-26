# Go Concept
## What is Go?
Go (also known as Golang) is a programming language developed at Google in 2007 and open-sourced in 2009.

It focuses on simplicity, reliability, and efficiency. It was designed to combine the efficacy, speed, and safety of a statically typed and compiled language with the ease of programming of a dynamic language to make programming more fun again.

In a way, they wanted to combine the best parts of Python and C++ so that they can build reliable systems that can take advantage of multi-core processors.


## Why Go 
* Easy to learn
* Fast and Reliable (Kubernetes and Docker )
* Simple yet powerful
* Career opportunities  

## What are modules?
A module is a collection of Go packages stored in a file tree with a go.mod file at its root, provided the directory is outside $GOPATH/src.

`$ go mod init example`

```module <name>

go <version>

require (
	...
)
```

`$ go install github.com/rs/zerolog`
`$ go list -m all`
`$ go mod tidy`
`go mod vendor`


## Packages 

A package is nothing but a directory containing one or more Go source files, or other Go packages.

This means every Go source file must belong to a package and package declaration is done at top of every source file as follows.
`package <package_name>`

## Workspaces

Workspaces allow us to work with multiple modules simultaneously without having to edit go.mod files for each module. Each module within a workspace is treated as a root module when resolving dependencies.

```shell
$ mkdir workspaces && cd workspaces
$ mkdir hello && cd hello
$ go mod init hello
```

## Other Useful commands 

`$ go fmt` : Starting with go fmt, which formats the source code and it's enforced by that language so that we can focus on how our code should work rather than how our code should look.
`$ go vet` : go vet which reports likely mistakes in our packages.
`go fix` finds Go programs that use old APIs and rewrites them to use newer ones.

`go generate` is usually used for code generation.
`go install` compiles and installs packages and dependencies
`go clean` is used for cleaning files that are generated by compile

`go build` : Compiles and Build a binary 
`go test` : Runs the go tests





## GOPATH vs GOROOT

GOPATH is a variable that defines the root of your workspace and it contains the following folders:
* src: contains Go source code organized in a hierarchy.
* pkg: contains compiled package code.
* bin: contains compiled binaries and executables.




https://www.karanpratapsingh.com/blog/learn-go-the-complete-course
https://go.dev/doc/effective_go
https://gobyexample.com/
https://tour.ardanlabs.com/tour/eng/list

## Interview Question [link](docs/Interview_Question.md)

## Channel

A channel is a communications pipe between goroutines. Things go in one end and come out another in the same order until the channel is closed.

## Goroutine [link](goroutine/README.md)
## Context
1. [Cancellation Singal with Cause](../context/cancellation_signal_with_cause.go)
2. [Context Deadline](../context/context_deadline.go)
3. [Cancellation Signal](../context/emit_cancellation_signal.go)
4. [Context Cancellation](../context/context_cancellation.go)
5. [Context Value](../context/context_values.go)

## Defer Keyword [link](defer/README.md)

## make and new [link](docs/make_and_new.md)

## Array [link](docs/arrays.md)

## Slices [link](docs/slices.md)

## Map [link](docs/map.md)

## Generic type in golang [link](generic_type/README.md)

## Embedding [link](https://tour.ardanlabs.com/tour/eng/embedding/1)

## Go Garbage Collection

1. https://medium.com/@souravchoudhary0306/exploring-the-inner-workings-of-garbage-collection-in-golang-tricolor-mark-and-sweep-e10eae164a12

2. https://medium.com/@souravchoudhary0306/under-the-hood-exploring-the-inner-workings-of-garbage-collection-in-golang-95b50bc8ce1d

3. https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8

## Go Goroutine Schdeuling 

1. https://medium.com/@sanilkhurana7/understanding-the-go-scheduler-and-looking-at-how-it-works-e431a6daacf


## Important Concepts

1. [decoupling](https://tour.ardanlabs.com/tour/eng/composition-decoupling/1)
2. [Interface Pollution](https://tour.ardanlabs.com/tour/eng/composition-pollution)


