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

	// use make to declare array if size is known
	// make array of int with size 10 but cap at 15
	j := make([]int, 10, 15)

	// load numbers from 1 to 10 and print
	for i := 0; i < len(j); i++ {
		j[i] = i + 1
	}
	fmt.Println("array:", j, "size:", len(j), "cap:", cap(j))

	// append until size 16 and print
	j = append(j, 11, 12, 13, 14, 15, 16)
	fmt.Println("array:", j, "size:", len(j), "cap:", cap(j))

	// multi-dimensional slice
	// declare 2 slices
	firstSlice := []string{"Naruto", "Luffy", "Natsu"}
	secondSlice := []string{"Simba", "Tmone", "Pumba"}
	fmt.Println("firstSlice:", firstSlice, "secondSlice:", secondSlice)

	allSlices := [][]string{firstSlice, secondSlice}
	fmt.Println("allSlices:", allSlices)

	// working with maps
	// declare map of string as index and int as value
	m := map[string]int{
		"Jett":      25,
		"Jelly":     22,
		"Frank":     70,
		"Piyawadee": 50,
	}
	// print the map
	fmt.Println("map:", m)

	// add a new key:value to the map and print
	m["Sam"] = 23
	fmt.Println("map:", m)

	// access map using comma ok idiom
	if v, ok := m["Jett"]; ok {
		fmt.Println(v, "is an awesome age! - Jett")
	}

	// access map using comma ok idiom for something not found
	if v, ok := m["John"]; ok {
		fmt.Println(v, "is an awesome age! - John")
	} else {
		fmt.Println("Couldn't find John")
	}

	// add John and print
	m["John"] = 32
	if v, ok := m["John"]; ok {
		fmt.Println(v, "is an awesome age! (after added) - John")
	}

	// use range to iterate through the whole map
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	// delete map using comma ok idiom
	if v, ok := m["John"]; ok {
		fmt.Println("Bye John whose", v)
		delete(m, "John")
	}

	// print map again
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

}
