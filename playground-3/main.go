package main

import (
	"fmt"
)

func main() {
	// working with for loops, Go doesn't have while loops
	for i := 0; i <= 10; i++ {
		fmt.Printf("Outer: %v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tInner: %v\n", j)
		}
	}

	// using for like while
	i := 0
	for i < 10 {
		fmt.Println("Use for to loop like while i:", i)
		i++
	}

	// using eternal for loop with break and continue
	i = 1
	for {
		i++
		// print until i > 7
		if i > 7 {
			break
		}
		if i%2 != 0 {
			// this skips when i is not an even number, goes to the next loop without next lines
			continue
		}
		fmt.Println("Eternal Loop printing i until > 7 then break:", i)
	}
	fmt.Println("Done!")

	// printing ascii using %#U
	// starting at 33 and ends at 122
	for j := 33; j <= 122; j++ {
		fmt.Printf("%v\t%#x\t%#U\n", j, j, j)
	}

	// using switch statement
	someNum := 7
	switch {
	case someNum == 0:
		fmt.Println("This is a zero")
	case someNum == 7:
		fmt.Println("This is a seven")
	case someNum%2 == 0:
		fmt.Println("This is an even number")
	case someNum%2 != 0:
		fmt.Println("This is an odd number")
	}

	// using switch statements with fallthrough and default
	fmt.Println("Using fallthrough and default")
	someNum2 := 7
	switch {
	case someNum2 == 0:
		fmt.Println("This is a zero")
		fallthrough
	case someNum2 == 7:
		fmt.Println("This is a seven")
		fallthrough
	case someNum2%2 == 0:
		fmt.Println("This is also an even number")
	case someNum2%2 != 0:
		fmt.Println("This is also an odd number")
	default:
		fmt.Println("This is default since nothing was true...")
	}

	// switching with values and multiple values
	someStr := "lit"
	switch someStr {
	case "lit", "dope", "awesome":
		// this is like using the or ||
		fmt.Println("This is fantabulous!")
	case "cool":
		fmt.Println("This is cool!")
	default:
		fmt.Println("This is default since nothing was true...")
	}
}
