package main

import "fmt"

type student struct {
	name string
}

func main() {
	s := student{"roushan"}
	s.setName()
	fmt.Println(s)
}

func (this *student) setName() {
	this.name = "kanchan"
}

// $ output:
// $ go run method.go
// {kanchan}
