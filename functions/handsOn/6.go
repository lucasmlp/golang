package main

func main() {

	x := 143
	y := 657
	sum := func(x, y int) int {
		return x + y
	}(x, y)
	println("The sum is", sum)
}
