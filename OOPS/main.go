package main

import (
	"fmt"
)

// Polymorphism
// func main() {
// 	var v Vechile

// 	v = NewBicycle()
// 	v.start()

// 	v = NewCar()
// 	v.start()
// }

// Inheritance
// func main() {
// 	C := NewCar()
// 	C.start()

// 	B := NewBicycle()
// 	B.start()
// }

// // Encapsulation
func main() {
	v := vechile.NewVechile(2, vechile.RED)
	fmt.Println(v.GetNoOfWheel())
}
