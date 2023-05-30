package main

import "fmt"

func main() {
	xs := []int{0, 1, 2, 3, 4}

	fooSum := foo(xs...)
	fmt.Println(fooSum)

	barSum := bar(xs)
	fmt.Println(barSum)
}

func foo(xs ...int) int {
	var sum int
	for _, x := range xs {
		sum += x
	}
	return sum
}

func bar(xs []int) int {
	var sum int
	for _, x := range xs {
		sum += x
	}
	return sum
}
