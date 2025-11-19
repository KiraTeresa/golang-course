package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	//info(c) // Cannot use 'c' (type circle) as the type shape Type does not implement 'shape' as the 'area' method has a pointer receiver
	fmt.Println(c.area())
}
