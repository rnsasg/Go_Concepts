package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr1 := make([]int, 5, 5)

	arr2 := new([5]int)[0:5]

	fmt.Println(arr1)
	fmt.Println(arr2)
	//make_vs_new_slice()

	//struct_example()
}

func make_vs_new_slice() {
	var p *[]int = new([]int)      // allocates slice structure; *p == nil; rarely useful
	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	fmt.Println(p) // &[]
	fmt.Println(v) // [0 0 0 0 0 0 0 0 0 0]

	// Unnecessarily complex:
	// var p *[]int = new([]int)
	// fmt.Println(p)
	*p = make([]int, 10)
	fmt.Println(p)

	// Idiomatic:
	// v := make([]int, 100)

}

func struct_example() {
	// Using new for a struct
	type Person struct {
		Name string
		Age  int
	}
	person := new(Person)

	fmt.Println(person)
	fmt.Println(person.Age)
	fmt.Println(unsafe.Sizeof(person.Name))
	fmt.Println(unsafe.Sizeof(person.Age))
}
