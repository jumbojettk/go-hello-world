package main

import (
	"fmt"
)

func main() {
	// ******************** EXERCISE 1 ********************
	// print every num from 1 to 10,000
	fmt.Println("******************** EXERCISE 1 ********************")

	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

	fmt.Println("******************** End EXERCISE 1 ********************")

	// ******************** EXERCISE 2 ********************
	// print rune code for all uppercase letter three times
	fmt.Println("******************** EXERCISE 2 ********************")

	for i2 := 65; i2 <= 90; i2++ {
		// print the num
		fmt.Println(i2)
		// print rune 3 times
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i2)
		}
	}

	fmt.Println("******************** End EXERCISE 2 ********************")

	// ******************** EXERCISE 3 ********************
	// print out all the years that was alive in the while form of for loop
	fmt.Println("******************** EXERCISE 3 ********************")

	year := 1993
	for year <= 2019 {
		fmt.Println(year)
		year++
	}

	fmt.Println("******************** End EXERCISE 3 ********************")

	// ******************** EXERCISE 4 ********************
	// use eternal for loop and break when condition met
	fmt.Println("******************** EXERCISE 4 ********************")

	year = 1993 // re-init year
	for {
		if year > 2019 {
			break
		}
		fmt.Println(year)
		year++
	}

	fmt.Println("******************** End EXERCISE 4 ********************")

	// ******************** EXERCISE 5 ********************
	// use mod to print out remainder of all num from 10 to 100 divide by 4
	fmt.Println("******************** EXERCISE 5 ********************")

	for num := 10; num <= 100; num++ {
		fmt.Printf("%v mod 4 = %v\n", num, num%4)
	}

	fmt.Println("******************** End EXERCISE 5 ********************")

	// ******************** EXERCISE 6 and 7 ********************
	fmt.Println("******************** EXERCISE 6 and 7 ********************")

	x := 2
	y := 4
	z := 7

	if (x + y + z) < 10 {
		fmt.Println("x + y + z is less than 10! It's", x+y+z)
	} else if (x + y + z) > 10 {
		fmt.Println("x + y + z is greater than 10! It's", x+y+z)
	} else {
		fmt.Println("x + y + z is equal to 10! It's", x+y+z)
	}

	fmt.Println("******************** End EXERCISE 6 and 7 ********************")

	// ******************** EXERCISE 8 ********************
	// use switch case without switch expression
	fmt.Println("******************** EXERCISE 8 ********************")

	switch {
	case false:
		fmt.Println("This aint gon print...")
	case true:
		fmt.Println("This will always print!")
	}

	fmt.Println("******************** End EXERCISE 8 ********************")

	// ******************** EXERCISE 9 ********************
	fmt.Println("******************** EXERCISE 9 ********************")

	favSport := "Football"

	switch favSport {
	case "Football":
		fmt.Println("Foot ball is my favorite sport!")
	case "Golf":
		fmt.Println("Golf is my favorite sport!")
	case "Basketball":
		fmt.Println("Basketball is my favorite sport!")
	case "Tennis":
		fmt.Println("Tennis is my favorite sport!")
	}

	fmt.Println("******************** End EXERCISE 9 ********************")

	// ******************** EXERCISE 10 ********************
	fmt.Println("******************** EXERCISE 10 ********************")

	// Write down what these print:
	fmt.Println(true && true, "should print true")
	fmt.Println(true && false, "should print false")
	fmt.Println(true || true, "should print true")
	fmt.Println(true || false, "should print true")
	fmt.Println(!true, "should print false")

	fmt.Println("******************** End EXERCISE 10 ********************")
}
