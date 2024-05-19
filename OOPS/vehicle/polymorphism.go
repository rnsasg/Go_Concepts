package vehicle

import "fmt"

type Vechile interface {
	start()
}

type Car struct {
}

type Bicycle struct {
}

func (c *Car) start() {
	fmt.Println("Start Car")
}

func (c *Bicycle) start() {
	fmt.Println("Start Bicycle")
}

func NewCar() Vechile {
	return &Car{}
}

func NewBicycle() Vechile {
	return &Bicycle{}
}
