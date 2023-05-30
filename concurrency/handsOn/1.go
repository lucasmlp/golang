package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("begin CPUs", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	numGoroutines := 2
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go printGoroutine(i)
	}

	fmt.Println("mid CPUs", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("end CPUs", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}

func printGoroutine(i int) {
	fmt.Println("goroutine", i)
	wg.Done()
}
