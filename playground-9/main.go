package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func init() {} can be used to declare things before main

// use sync.WaitGroup to tell main() to wait for the Goroutines to finish before exiting
var wg sync.WaitGroup

func main() {
	// WORKING WITH WAITGROUPS
	// These are contants
	fmt.Println("OS\t", runtime.GOOS)     // Print the operating system of the machine
	fmt.Println("ARCH\t", runtime.GOARCH) // Print the architecture of the machine

	// These are functions
	fmt.Println("CPUs\t", runtime.NumCPU())             // Print the number of CPU on machine
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // Print the number of Goroutines running

	wg.Add(1)
	go foo()
	bar()

	// Print again after starting more goroutines
	fmt.Println("CPUs\t", runtime.NumCPU())             // Print the number of CPU on machine
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // Print the number of Goroutines running
	wg.Wait()

	// WORKING WITH CHANNELS
	ch := make(chan int) // integer channel for receiving int outside of Goroutine
	go func() {
		ch <- multByFive(5) // assign returned value from func to channel
	}() // function literal must have ending prenthesis to invoke
	fmt.Println("Working with channels, result of 5 multiplied by 5 is:", <-ch)
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func multByFive(num int) int {
	return num * 5
}
