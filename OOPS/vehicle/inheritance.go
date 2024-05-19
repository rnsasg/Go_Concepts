package vehicle

import "fmt"

type Vechile struct {
	wheel int
}

func (v *Vechile) start() {
	fmt.Println("Starting a vehicle with %d wheel", v.wheel)
}

// Composition
type Car struct {
	Vechile
}

// Composition
type Bicycle struct {
	Vechile
}

func (c *Car) start() {
	fmt.Println("Start Car")
}

func (c *Bicycle) start() {
	fmt.Println("Start Bicycle")
}

func NewCar() *Car {
	return &Car{Vechile{wheel: 4}}
}

func NewBicycle() *Bicycle {
	return &Bicycle{Vechile{wheel: 2}}
}

func main() {

}
