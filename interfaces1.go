package main

import "fmt"

type tank interface {
	area() float64
	volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

func (m myvalue) area() float64 {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myvalue) volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank
	t = myvalue{19, 28}

	fmt.Println("area of tank is:", t.area())
	fmt.Println("volume of tank is:", t.volume())
}
