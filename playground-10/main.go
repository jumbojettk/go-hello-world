package main

import "fmt"

func main() {
	// UNDERSTANDING CHANNELS
	c := make(chan int)

	// Channels are blocking, it needs its own routine
	// c <- 7	// this will not work
	go func() {
		c <- 7 // this will work since channels are blocking since this routine gets fired off to assign 7 to the channel while main() is running
	}()

	fmt.Println("This is the result from using a Goroutine with a channel: ", <-c)

	// Channels can have buffers, so that it can store values without its own Goroutine
	// but why use this if you can just assign variables, this shouldn't be practiced, but it's available
	c2 := make(chan int, 2) // 2 buffer spaces for channel to store val

	c2 <- 7  // store 7 as first val
	c2 <- 28 // store 28 as second val

	// Now c2 has 2 vals stored, can be read from FIFO (first to last)
	fmt.Println("This is the result from using a channel buffer with a channel, value 1: ", <-c2)
	fmt.Println("This is the result from using a channel buffer with a channel, value 2: ", <-c2)
}
