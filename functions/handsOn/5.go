package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

type square struct {
	l float64
	w float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (s square) area() float64 {
	return s.l * s.w
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("area:", s.area())
}

func main() {
	c1 := circle{
		12.345,
	}

	s1 := square{
		15,
		15,
	}

	info(c1)
	info(s1)
}
