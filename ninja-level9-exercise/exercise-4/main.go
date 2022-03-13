package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("******************** EXERCISE 4 ********************")
	// Fix the race condition you created in the previous exercise by using a mutex
	// it makes sense to remove runtime.Gosched()
	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of GoRoutines:", runtime.NumGoroutine())

	incrementer := 0 // start at 0

	const maxgs = 100     // max go routines
	var wg sync.WaitGroup // wait group
	wg.Add(maxgs)

	var mu sync.Mutex // mutex

	// loop a bunch of routines
	for i := 0; i < maxgs; i++ {
		go func() {
			mu.Lock()             // lock the mutex
			incVal := incrementer // read
			incVal++              //	inc
			incrementer = incVal  // set new val
			mu.Unlock()           // unlock when done
			wg.Done()             // this is where the wg gets done
		}()

	}

	// use waitgroups to wait for all of your goroutines to finish
	wg.Wait()                                            // wait for all routines to finish
	fmt.Println("Final Incrementer Value:", incrementer) // final value

	// the above will create a race condition.
	// Prove that it is a race condition by using the -race flag
}
