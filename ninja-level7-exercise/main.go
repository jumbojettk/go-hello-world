package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	// ******************* EXERCISE 1 *******************
	// Create a value and assign it to a variable.
	goodNum := 7

	// Print the address of that value.
	fmt.Println("Good Num:", goodNum)
	fmt.Println("Good Num Addy:", &goodNum)

	// ******************* EXERCISE 2 *******************
	// create a value of type person
	sillyDude := person{
		first: "Jett",
		last:  "Krassner",
		age:   25,
	}

	// print out the value
	fmt.Println("Silly Dude:", sillyDude)

	// call “changeMe”
	changeMe(&sillyDude)

	// print out the value
	fmt.Println("Changed Silly Dude:", sillyDude)

	// call "changeMeh" for kicks using the (*p)
	changeMeh(&sillyDude)

	// print out the value
	fmt.Println("Changed-Meh Silly Dude:", sillyDude)

}

func changeMe(p *person) {
	p.first = "Jettster" // same as using (*p).first
	p.last = "Krassnerski"
	p.age = 25000
}

func changeMeh(p *person) {
	(*p).first = "Jettie"
	(*p).last = "Krassnie"
	(*p).age = 2500
}
