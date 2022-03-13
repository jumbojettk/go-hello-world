package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("******************** EXERCISE 3 ********************")
	// Using goroutines, create an incrementer program
	// have a variable to hold the incrementer value
	// launch a bunch of goroutines
	// each goroutine should
	// read the incrementer value
	// store it in a new variable
	// yield the processor with runtime.Gosched()
	// increment the new variable
	// write the value in the new variable back to the incrementer variable
	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of GoRoutines:", runtime.NumGoroutine())

	incrementer := 0 // start at 0

	const maxgs = 100     // max go routines
	var wg sync.WaitGroup // wait group
	wg.Add(maxgs)

	// loop a bunch of routines
	for i := 0; i < maxgs; i++ {
		go func() {
			fmt.Println("Incrementer val: ", incrementer)
			incVal := incrementer // read
			runtime.Gosched()     // yield the processor (hold up)
			incVal++              //	inc
			incrementer = incVal  // set new val
			wg.Done()             // this is where the wg gets done
		}()

	}

	// use waitgroups to wait for all of your goroutines to finish
	wg.Wait()                                            // wait for all routines to finish
	fmt.Println("Final Incrementer Value:", incrementer) // final value

	// the above will create a race condition.
	// Prove that it is a race condition by using the -race flag
}
