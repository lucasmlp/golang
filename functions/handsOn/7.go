package main

import "fmt"

func main() {

	x := 143
	y := 657
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Printf("The type of sum is %T\n", sum)
	println("The sum of", x, "and", y, "is", sum(x, y))
}
