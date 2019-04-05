package main

import (
	"fmt"
)

const (
	pi = 3.1415
)

// struct with identifier person, or person type struct
// the fields: first, last, age
type person struct {
	first string
	last  string
	age   int
}

// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
func (p person) speak() {
	fmt.Println("Hi! My name is", p.first, p.last, ". I'm", p.age, "! :)")
}

// create a type SQUARE
type square struct {
	length float64
	width  float64
}

// create a type CIRCLE
type circle struct {
	radius float64
}

// attach a method to each that calculates AREA and returns it
// square area = L * W
func (sqr square) area() float64 {
	return sqr.length * sqr.width
}

// circle area= π r 2
func (cir circle) area() float64 {
	return pi * (cir.radius * cir.radius)
}

// create a type SHAPE that defines an interface as anything that has the AREA method
type shape interface {
	area() float64
}

// create a func INFO which takes type shape and then prints the area
func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	fmt.Println("******************** EXERCISE 1 ********************")
	// create a func with the identifier foo that returns an int
	// create a func with the identifier bar that returns an int and a string
	// call both funcs
	// print out their results
	fmt.Println(foo())
	fmt.Println(bar()) // can print multiple values

	fmt.Println("******************** EXERCISE 2 ********************")
	// create a func with the identifier foo2 that
	// takes in a variadic parameter of type int
	// pass in a value of type []int into your func (unfurl the []int)
	// returns the sum of all values of type int passed in
	xi2 := []int{2, 3, 4, 5, 6, 7}
	x2 := foo2(xi2...)
	fmt.Println(x2)

	// create a func with the identifier bar2 that
	// takes in a parameter of type []int
	// returns the sum of all values of type int passed in
	x2 = bar2(xi2)
	fmt.Println(x2)

	fmt.Println("******************** EXERCISE 3 ********************")
	// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	startProcess()

	fmt.Println("******************** EXERCISE 4 ********************")
	// create a value of type person
	p1 := person{
		first: "Jett",
		last:  "Krassner",
		age:   25,
	}
	fmt.Println(p1)

	// call the method from the value of type person
	p1.speak()

	fmt.Println("******************** EXERCISE 5 ********************")
	// create a value of type square
	sqr1 := square{
		length: 7,
		width:  7,
	}

	// create a value of type circle
	cir1 := circle{
		radius: 7,
	}

	// use func info to print the area of square
	info(sqr1)

	// use func info to print the area of circle
	info(cir1)

	fmt.Println("******************** EXERCISE 6 ********************")
	// Build and use an anonymous func
	func(s string) {
		fmt.Println("Anonymously", s)
	}("awesome!")

	fmt.Println("******************** EXERCISE 7 ********************")
	// Assign a func to a variable, then call that func
	f7 := func(s string) {
		fmt.Println("This expression is", s)
	}

	f7("so cool!")

	fmt.Println("******************** EXERCISE 8 ********************")
	// assign the returned func to a variable
	f8 := returnFunc()

	// call the returned func
	s8 := f8()
	fmt.Println(s8) // should return string

	fmt.Println("******************** EXERCISE 9 ********************")
	// Callback, pass a func into a func as an argument
	sName, sCoolnm := findInfo(coolName, p1)
	fmt.Println(sName, "to", sCoolnm)

	fmt.Println("******************** EXERCISE 10 ********************")
	// create a func which “encloses” the scope of a variable:
	f10 := enclosed()
	fmt.Println(f10())
	fmt.Println(f10())
	fmt.Println(f10())
	fmt.Println(f10())
	f10too := enclosed()
	fmt.Println(f10too())
	fmt.Println(f10too())
	fmt.Println(f10too())

}

// func foo returns int
func foo() int {
	return 7
}

// func bar returns int and string
func bar() (int, string) {
	return 25, "Jettster"
}

// func foo2 takes variadic int
// returns sum of all value
func foo2(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// func bar2 takes slice of int
// returns sum of all value
func bar2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// func open
func open() {
	fmt.Println("Opening...")
}

// func close
func close() {
	fmt.Println("Closing...")
}

// func startProcess
func startProcess() {
	defer close()
	open()

	fmt.Println("Starting process from opening a file...")
	fmt.Println("Processing file for ninjaness...")
	fmt.Println("NINJA!!!")
	// deferred close will invoke here
}

// func that returns a func(that returns a string)
func returnFunc() func() string {
	return func() string {
		return "Woo! This is cool!"
	}
}

// func name that takes fn and ln strings
// returns a firstname and lastname as one string
func coolName(fn string, ln string) string {
	fn += "ster"
	ln += "ski"
	return fn + " " + ln
}

// func that takes type person struct
func findInfo(f func(fn string, ln string) string, p person) (string, string) {
	name := p.first + " " + p.last
	coolNm := f(p.first, p.last)
	return name, coolNm
}

// func that encloses a scope of var
func enclosed() func() int {
	x := 0 // declare x out here to limit scope
	return func() int {
		x++
		return x
	}
}
