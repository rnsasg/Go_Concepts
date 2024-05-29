package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Using make for a slice
	s := make([]int, 5)
	fmt.Println("Slice:", s) // Output: Slice: [0 0 0 0 0]

	// Using new for an integer
	p := new(int)
	fmt.Println("Pointer to int:", *p) // Output: Pointer to int: 0
	*p = 42
	fmt.Println("Pointer to int after assignment:", *p) // Output: Pointer to int after assignment: 42

	// Using make for a map
	m := make(map[string]int)
	m["key"] = 1
	fmt.Println("Map:", m) // Output: Map: map[key:1]

	// Using new for a struct
	type Person struct {
		Name string
		Age  int
	}
	person := new(Person)

	// fmt.Println(person)
	// fmt.Println(person.Age)
	fmt.Println(unsafe.Sizeof(person.Name))
	fmt.Println(unsafe.Sizeof(person.Age))

	//fmt.Println()

	person.Name = "Alice"
	person.Age = 30
	fmt.Println("Person:", *person) // Output: Person: {Alice 30}
}

func make_vs_new_slice() {
	var p *[]int = new([]int)      // allocates slice structure; *p == nil; rarely useful
	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	// Unnecessarily complex:
	var p *[]int = new([]int)
	*p = make([]int, 100, 100)

	// Idiomatic:
	v := make([]int, 100)

}
