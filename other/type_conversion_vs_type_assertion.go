package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Type Conversion

	i := 42
	f := float64(i)
	u := uint(f)

	//fmt.Printf("%T %T", f, u)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(u))

	// Type Assertion
	var j interface{} = "my string"
	s := j.(string)
	fmt.Printf("%s", s)

}
