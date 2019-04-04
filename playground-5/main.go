package main

import (
	"fmt"
)

// declare a type struct
type person struct {
	first string
	last  string
	age   int
}

// declare another type struct and embed person
type broski struct {
	person
	lovesCoding bool
	lovesBeer   bool
}

func main() {
	// declare value of type person
	p1 := person{
		first: "Jett",
		last:  "Krassner",
		age:   25,
	}

	fmt.Println("p1:", p1)

	// accessing fields value in type person
	fmt.Printf(" First: %v\n Last: %v\n Age: %v\n Type: %T\n", p1.first, p1.last, p1.age, p1)

	// declare broski and embed person
	bro1 := broski{
		person: person{
			first: "Jettster",
			last:  "Krassnerski",
			age:   25,
		},
		lovesCoding: true,
		lovesBeer:   true,
	}

	// accessing fields value in embedded type (fields are promoted)
	fmt.Printf(" First: %v\n Last: %v\n Age: %v\n Type: %T\n", bro1.first, bro1.last, bro1.age, bro1)

	// using anonymouse type struct
	fam := struct {
		last string
		size int
	}{
		last: "Krassner",
		size: 4,
	}
	// print anonymouse type struct
	fmt.Printf(" Last: %v\n Size: %v\n Type: %T\n", fam.last, fam.size, fam)

}
