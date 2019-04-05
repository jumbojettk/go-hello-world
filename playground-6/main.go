package main

import (
	"fmt"
)

// STRUCTS
// declare type struct
type person struct {
	first string
	last  string
}

// declare person embedded type struct
type ninja struct {
	person
	skills []string
}

// declare something unrelated to person
type dog struct {
	name  string
	breed string
}

// INTERFACES
// create an interface for human
// any type that speaks can use functions for human
type human interface {
	speak()
}

// create an interface for all
type fromSeattle interface {
}

// TYPE METHODS
// create method for person
func (p person) speak() {
	fmt.Println("Hi! I'm", p.first, p.last, ". :)")
}

// create method for ninja
func (n ninja) disappear() {
	fmt.Println(n.first, n.last, "disappeared!... Where'd he go?!!! o_O!!!")
}

// method for human
func smile(h human) {
	fmt.Println("I'm smiling!", h)
}

// method for human with assertion on types
func hug(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I'm a person who loves a hug! - ", h.(person).first, h.(person).last)
	case ninja:
		fmt.Println("Like my ninjaness? *Hugs! - ", h.(ninja).first, h.(ninja).last)
	}

	fmt.Println("Hugs are awesome!", h)
}

// method for dog
func (d dog) bark() {
	fmt.Println("Woof woof, woof woof woof! ^.^ - ", d.name)
}

// method for fromSeattle
func seattle(f fromSeattle) {
	fmt.Println("I'm from Seattle!", f)
}

func main() {
	// Invoke the functions
	s1 := "This is so cool!"
	printString(s1)
	sToPrint := fmt.Sprintf("The size of s1 is: %v", getSize(s1))
	printString(sToPrint)

	size1, s4 := getSizeAndFirstFour(s1)
	sToPrint = fmt.Sprintf("The size of s1 is: %v\nThe first 4 letters of s1 is: %v", size1, s4)
	printString(sToPrint)

	// Invoke variadic functions
	printUnlimitInt("random numbers", 2, 4, 7, 27, 77)

	// variadic functions with slices
	// use unfurling
	xi := []int{22, 55, 77, 99, 1019}
	printUnlimitInt("random number in slice", xi...)

	// use defer in a function
	printLast("This is so cool!", 7, 77)

	// USING METHODS FOR TYPES
	// declare new person type and assign values
	p1 := person{
		first: "Jett",
		last:  "Krassner",
	}

	// declare new ninja type
	n1 := ninja{
		person: person{
			first: "Jettster",
			last:  "Krassnerski",
		},
		skills: []string{"coding", "problem solving", "initiative"},
	}

	// declare new dog type
	d1 := dog{
		name:  "Simba",
		breed: "Labrador Retriever",
	}

	// use person method
	p1.speak()

	// use ninja method
	n1.speak() // since ninja is also a person
	n1.disappear()

	// use human interface
	smile(p1)
	smile(n1)

	// use another human interface with type assertion
	hug(p1)
	hug(n1)

	// use dog method
	// person or ninja cannot use dog methods
	d1.bark()

	// use fromSeattle interface, open to all
	seattle(p1)
	seattle(n1)
	seattle(d1)

	// using anonymous function
	func(s string) {
		fmt.Println("Anonymously,", s)
	}("awesome!")

	// using function expression
	f1 := func(s string) {
		fmt.Println("This expression is", s)
	}

	f1("so cool!")

	// returning a function
	f2 := dope()
	fmt.Printf("This is the returned function's type: %T\n", f2)
	x2 := f2()
	fmt.Println("Invoking returned function, should be 7", x2)

	// callback function
	xi1 := []int{2, 3, 4, 7, 7, 8, 10, 34}
	totalSum := sum(xi1...)
	fmt.Println(totalSum)
	totalEvenSum := evenSum(sum, xi1...)
	fmt.Println(totalEvenSum)

	// using recursive function
	fmt.Println("Recusive factorial:", factorial(7)) // should be 5040

	// using loop to do the same function as recursive func
	fmt.Println("Loop factorial:", loopFactorial(7))
}

// Takes a string to print
// returns nothing
func printString(s string) {
	fmt.Println(s)
}

// takes a string, find the size
// returns the int of size
func getSize(s string) int {
	return len(s)
}

// takes a string, find the size
// returns the first 4 letters
func getSizeAndFirstFour(s string) (int, string) {
	return len(s), s[:4]
}

// takes a string and series of ints, prints out the string
// print all ints
func printUnlimitInt(s string, x ...int) {
	fmt.Println(s)
	for i, v := range x {
		fmt.Println("\t", i, v)
	}
}

// defer printing to the last thing
func printLast(s string, x int, y int) {
	defer fmt.Println("y:", y)      // this will be last
	defer fmt.Println("x:", x)      // this be second to last
	defer fmt.Println("String:", s) // this be third to last

	s = s + " Cuz it's awesome!"
	x += x
	y *= y

	fmt.Println("String:", s)
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("But these are the previous values")
	// the deferred lines will execute after this one
}

// a dope function that returns func(that returns int)
func dope() func() int {
	return func() int {
		return 7
	}
}

// takes a variadic int as a slice and find then return the sum
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

func evenSum(f func(xi ...int) int, xi2 ...int) int {
	xiEven := []int{} // declare slice to store even num

	for _, v := range xi2 {
		// determin even nums and append to slice
		if v%2 == 0 {
			xiEven = append(xiEven, v)
		}
	}

	// call the function on the even slice
	total := f(xiEven...)

	return total
}

// recursive func for factorial
func factorial(n int) int {
	// base case
	if n == 0 {
		return 1
	}

	// return the n * n-1 * n-2 * .... * 1
	return n * factorial(n-1)
}

// loopy func for factorial
func loopFactorial(n int) int {
	total := 1 // 1 x anything = 1

	for n > 0 {
		total *= n
		n--
	}

	return total
}
