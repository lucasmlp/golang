package main

import "fmt"

func main() {
	defer fmt.Println("Defering main")

	xs := []int{0, 1, 2, 3, 4}
	var sum int

	for _, x := range xs {
		fmt.Println(x)
		sum += x
	}

	fmt.Println(sum)
}
