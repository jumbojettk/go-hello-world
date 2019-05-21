package main

import (
	"fmt"
)

func main() {
	// declare variables
	num1 := 7
	ptNum := &num1

	fmt.Println("Number: ", num1)
	fmt.Println("Address: ", ptNum)
	fmt.Println("Dereference the address: ", *ptNum)
	fmt.Printf("Num type: %T\nPointer type: %T\nDereference Type: %T\n", num1, ptNum, *ptNum)

	// now set the value of num1 to a different value using the pointer
	fmt.Println("Making some changes......................")
	*ptNum = 77
	fmt.Println("New Number: ", num1)
	fmt.Println("Same Address: ", ptNum)
	fmt.Println("Dereference the address: ", *ptNum)
	fmt.Printf("Num type: %T\nPointer type: %T\nDereference Type: %T\n", num1, ptNum, *ptNum)

}
