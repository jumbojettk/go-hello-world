package main

import (
	"fmt"
)

const (
	a     = 7
	b int = 8
)

const (
	_          = iota
	nextYear   = 2019 + iota
	next2Years = 2019 + iota
	next3Years = 2019 + iota
	next4Years = 2019 + iota
)

func main() {
	// ******************** EXERCISE 1 ********************
	// print a number in dec, bin, hex
	fmt.Println("******************** EXERCISE 1 ********************")
	num := 7
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Binary: %b\n", num)
	fmt.Printf("Hex: %#X\n", num)

	fmt.Println("******************** End EXERCISE 1 ********************")

	// ******************** EXERCISE 2 ********************
	fmt.Println("******************** EXERCISE 2 ********************")
	// all will return true
	equals := 2 == 2
	lessEqual := 4 <= 7
	greatEqual := 7 >= 2
	notEqual := 6 != 7
	lesser := 2 < 7
	greater := 7 > 1
	fmt.Print("Printing result of expression from the operators: ")
	fmt.Println(equals, lessEqual, greatEqual, notEqual, lesser, greater)

	fmt.Println("******************** End EXERCISE 2 ********************")

	// ******************** EXERCISE 3 ********************
	// create typed and untyped const
	fmt.Println("******************** EXERCISE 3 ********************")
	fmt.Print("Printing result of typed and untyped const: ")
	fmt.Println(a, b)
	fmt.Println("******************** End EXERCISE 3 ********************")

	// ******************** EXERCISE 4 ********************
	// assign int to variable and print in dec, bin, hex
	// shift 1 bit to left and assign to variable then print in dec, bin, hex
	fmt.Println("******************** EXERCISE 4 ********************")
	shiftInt := 7
	fmt.Printf("dec: %d,\tbin: %b,\thex: %#X\n", shiftInt, shiftInt, shiftInt)

	shiftedInt := shiftInt << 1
	fmt.Printf("Shifted! dec: %d,\tbin: %b,\thex: %#X\n", shiftedInt, shiftedInt, shiftedInt)

	fmt.Println("******************** End EXERCISE 4 ********************")
	// ******************** EXERCISE 5 ********************
	// use string literal to print
	fmt.Println("******************** EXERCISE 5 ********************")
	fmt.Println(
		`This is a string literal with "quotes and stuff"
	
	this is a couple of new lines without newline escape
	`)

	fmt.Println("******************** End EXERCISE 5 ********************")
	// ******************** EXERCISE 6 ********************
	fmt.Println("******************** EXERCISE 6 ********************")
	fmt.Printf("Next Year: %v\nNext 2 Years: %v\nNext 3 Years: %v\nNext 4 Years: %v\n", nextYear, next2Years, next3Years, next4Years)

	fmt.Println("******************** End EXERCISE 6 ********************")
}
