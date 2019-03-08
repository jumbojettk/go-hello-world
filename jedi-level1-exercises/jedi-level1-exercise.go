package main

import (
	"fmt"
)

// for Exercise 2
var a int
var b string
var c bool

// for Exercise 4
type awesome int

var j awesome

// for Exercise 5
// underlying type of custom type j
var k int

func main() {
	// ****************** EXERCISE 1 ******************
	// Declare and assign values
	fmt.Println("******** Exercise 1 ********")
	x := 42
	y := "James Bond"
	z := true

	// single print
	fmt.Printf("x = %v\ny = %v\nz = %v\n", x, y, z)

	// multiple print
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)

	fmt.Println("******** End Exercise 1 ********")

	// ****************** EXERCISE 2 ******************
	// Declare vars without assigning values in the package level scope
	// these vars will have zero values (default) assigned to them
	fmt.Println("******** Exercise 2 ********")

	// print all values
	fmt.Printf("a = %v\nb = %v\nc = %v\n", a, b, c)

	fmt.Println("******** End Exercise 2 ********")

	// ****************** EXERCISE 3 ******************
	fmt.Println("******** Exercise 3 ********")
	a = 42
	b = "James Bond"
	c = true

	// use Sprintf print all values as string and assign to variable to print and print it
	strPrint := fmt.Sprintf("a = %v\nb = %v\nc = %v\n", a, b, c)
	fmt.Print(strPrint)

	fmt.Println("******** End Exercise 3 ********")

	// ****************** EXERCISE 4 ******************
	fmt.Println("******** Exercise 4 ********")
	// print j value and type
	fmt.Printf("value of j = %v\ntype of j = %T\n", j, j)

	// assign j to 42 using conversion
	j = 42
	fmt.Printf("value of j = %v\ntype of j = %T\n", j, j)
	fmt.Println("******** End Exercise 4 ********")

	// ****************** EXERCISE 5 ******************
	fmt.Println("******** Exercise 5 ********")
	// print j value and type
	fmt.Printf("value of k = %v\ntype of k = %T\n", k, k)

	// assign j to 42 using conversion
	k = int(j)
	fmt.Printf("value of k = %v\ntype of k = %T\n", k, k)
	fmt.Println("******** End Exercise 5 ********")
}
