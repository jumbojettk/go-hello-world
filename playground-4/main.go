package main

import (
	"fmt"
)

func main() {
	// declare an array with size and type
	var arr1 [5]int

	// assign some values to array index
	arr1[2] = 7
	arr1[4] = 7

	fmt.Println(arr1)

	// composite literal, slice of ints
	x := []int{4, 5, 6, 7, 8, 77}
	fmt.Println(x)

	// print them by index
	fmt.Println(x[5])

	// print by looping index and value using range
	for i, v := range x {
		fmt.Println(i, v)
	}

	// slicing a slice
	fmt.Println(x[1:])  // print from index 1 to the end
	fmt.Println(x[1:4]) // print from index 1 to 3 not including 4

	// print using regular for loop
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	// append to a slice
	x = append(x, 8, 8, 28, 29, 7)
	fmt.Println(x)

	// append a slice to a slice, note the "..."
	y := []int{2, 3, 7, 89, 200, 256, 1000}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)

	// deleting values off a slice using append
	// take values from 0 to 3 and append values from 5 to 10
	// this will leave index 4 out of the list and only goes up to index 10
	x = append(x[:4], x[5:11]...)
	fmt.Println(x)

}
