package main

import (
	"fmt"
)

// working with iota
const (
	a = iota
	b
	c
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	// Working with type String
	str := "This is just a normal string, ya feel?"
	fmt.Printf("String Value: %v\nString Type: %T\n", str, str)

	// loading string into array/slice of bytes and assign to new var
	bStr := []byte(str)
	fmt.Printf("Byte String Value: %v\nByte String Type: %T\n", bStr, bStr) // print it
	fmt.Printf("Byte of first char in bStr in decimal: %d\n", bStr[0])
	fmt.Printf("Byte of first char in bStr in binary: %b\n", bStr[0])
	fmt.Printf("Byte of first char in bStr in hex: %#X\n", bStr[0])

	// working with iota
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// working with bit shifting
	num := 4
	fmt.Printf("Value: %v, Binary: %b\n", num, num)

	num2 := num << 1 // shift left by 1 (adds an extra 0 at the end)
	fmt.Printf("Shifted left by 1 value: %v, Binary: %b\n", num2, num2)

	// comment these out and use const with iota and shifting
	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
