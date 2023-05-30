package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and I'm", p.age, "years old")
}

func main() {
	p1 := person{
		"Lucas",
		"Machado",
		32,
	}

	p1.speak()
}
