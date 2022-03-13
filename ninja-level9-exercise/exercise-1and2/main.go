package main

import (
	"fmt"
	"runtime"
	"sync"
)

// create a type person struct
// attach a method speak to type person using a pointer receiver
// *person
// create a type human interface
// to implicitly implement the interface, a human must have the speak method
// create func “saySomething”
// have it take in a human as a parameter
// have it call the speak method
// show the following in your code
// you CAN pass a value of type *person into saySomething
// you CANNOT pass a value of type person into saySomething
type person struct {
	FirstName string
	LastName  string
	Age       uint
}

func (p *person) speak() {
	fmt.Printf("Sup! I'm %v, and I'm %v! :)\n", p.FirstName, p.Age)
}

type human interface {
	speak()
}

func main() {
	fmt.Println("******************** EXERCISE 1 ********************")
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(4) // in addition to the main goroutine, launch two additional goroutines

	go func() {
		fmt.Println("hello from thing one")
		fmt.Println("gs", runtime.NumGoroutine())
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing two")
		fmt.Println("gs", runtime.NumGoroutine())
		wg.Done()
	}()

	// each additional goroutine should print something out
	go func() {
		fmt.Println("uh suh from thing tres")
		fmt.Println("gs", runtime.NumGoroutine())
		wg.Done() // use waitgroups to make sure each goroutine finishes before your program exists
	}()

	go func() {
		fmt.Println("sup from thing quatro")
		fmt.Println("gs", runtime.NumGoroutine())
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

	fmt.Println("******************** EXERCISE 2 ********************")
	p1 := person{
		FirstName: "Jett",
		LastName:  "Krassner",
		Age:       25,
	}

	//saySomething(p1)  // cannot pass a person
	saySomething(&p1) // can pass a person's mem address (*person)

	p2 := &person{
		FirstName: "Jettster",
		LastName:  "Krassnerski",
		Age:       25,
	}

	var p3 human
	p3 = p2

	saySomething(p3) // can pass *person type
	saySomething(p2) // can pass *person type

}

// saySomething
func saySomething(h human) {
	h.speak()
}
