package main

import (
	"fmt"
)

func main() {
	fmt.Println("******************** EXERCISE 1 ********************")
	// use composite literal
	// create array with 5 ints
	arr1 := [5]int{11, 22, 33, 44, 55}

	// range and print out values
	for i, v := range arr1 {
		fmt.Println(i, ":", v)
	}

	// use printf to print out Type of array
	fmt.Printf("Type of array: %T\n", arr1)

	fmt.Println("******************** EXERCISE 2 ********************")
	// Using a COMPOSITE LITERAL:
	// create a SLICE of TYPE int, // assign 10 VALUES
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Range over the slice and print the values out.
	for i, v := range xi {
		fmt.Println(i, ":", v)
	}

	// Using format printing
	// print out the TYPE of the slice
	fmt.Printf("Type of slice: %T\n", xi)

	fmt.Println("******************** EXERCISE 3 ********************")
	// Using the code from the previous example,
	// use SLICING to create the following new slices which are then printed:
	// [42 43 44 45 46]
	xi1 := xi[:5]
	fmt.Println("xi1:", xi1)

	// [47 48 49 50 51]
	xi2 := xi[5:]
	fmt.Println("xi2:", xi2)

	// [44 45 46 47 48]
	xi3 := xi[2:7]
	fmt.Println("xi3:", xi3)

	// [43 44 45 46 47]
	xi4 := xi[1:6]
	fmt.Println("xi4:", xi4)

	fmt.Println("******************** EXERCISE 4 ********************")
	// start with this slice
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)

	// append to that slice this value: 52
	// print out the slice
	xi = append(xi, 52)
	fmt.Println(xi)

	// in ONE STATEMENT append to that slice these values: 53, 54, 55
	// print out the slice
	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)

	// append to the slice this slice
	// y := []int{56, 57, 58, 59, 60}
	// print out the slice
	y := []int{56, 57, 58, 59, 60}
	xi = append(xi, y...)
	fmt.Println(xi)

	fmt.Println("******************** EXERCISE 5 ********************")
	// 	start with this slice
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xi = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// use APPEND & SLICING to get these values here
	// which you should ASSIGN to a variable “y” and then print:
	// [42, 43, 44, 48, 49, 50, 51]
	y = append(xi[:3], xi[6:]...)
	fmt.Println("xi:", xi, "\ny:", y)

	fmt.Println("******************** EXERCISE 6 ********************")
	// Create a slice to store the names of all of the states in the United States of America.
	xs := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`,
		` California`, ` Colorado`, ` Connecticut`, ` Delaware`,
		` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`,
		` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
		` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`,
		` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`,
		` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
		` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
		` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}

	// What is the length of your slice?
	fmt.Println("Length of slice:", len(xs))

	// What is the capacity?
	fmt.Println("Cap of slice:", cap(xs))

	// Print out all of the values,
	// along with their index position in the slice,
	// without using the range clause.
	for i := 0; i < len(xs); i++ {
		fmt.Println(i, xs[i])
	}

	fmt.Println("******************** EXERCISE 7 ********************")
	// 	Create a slice of a slice of string ([][]string).
	// Store the following data in the multi-dimensional slice:
	// "James", "Bond", "Shaken, not stirred"
	// "Miss", "Moneypenny", "Helloooooo, James."
	xs2D := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	// Range over the records, then range over the data in each record.
	for i, v := range xs2D {
		fmt.Println("Slice:", i)
		for i, v2 := range v {
			fmt.Println("\tIndex:", i, "\tValue:", v2)
		}
	}

	fmt.Println("******************** EXERCISE 8 ********************")
	// 	Create a map with a key of TYPE string which is a person’s “last_first” name,
	// and a value of TYPE []string which stores their favorite things.
	// Store three records in your map.
	// 	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	// Print out all of the values, along with their index position in the slice.
	for k, xFavThings := range m {
		fmt.Println(k)

		for i, favThing := range xFavThings {
			fmt.Printf("\tIndex: %v,\tFavorite Thing: %v\n", i, favThing)
		}
	}

	fmt.Println("******************** EXERCISE 9 ********************")
	// Using the code from the previous example,
	// add a record to your map.
	m[`krassner_jett`] = []string{`Coding`, `Being awesome`, `Trying new things`}

	// Now print the map out using the “range” loop
	for k, xFavThings := range m {
		fmt.Println(k)

		for i, favThing := range xFavThings {
			fmt.Printf("\tIndex: %v,\tFavorite Thing: %v\n", i, favThing)
		}
	}

	fmt.Println("******************** EXERCISE 10 ********************")
	// Using the code from the previous example,
	// delete a record from your map.
	delete(m, "no_dr")

	// Now print the map out using the “range” loop
	for k, xFavThings := range m {
		fmt.Println(k)

		for i, favThing := range xFavThings {
			fmt.Printf("\tIndex: %v,\tFavorite Thing: %v\n", i, favThing)
		}
	}

}
