package main

// import pkgs
import "fmt"

func main() {
	fmt.Println("Uh Suhhh!")
	lit()

	fmt.Println("Bouta print \"lit\" during \"even\" and running 77 times!")

	// notice the parenthesis not used, no need
	// stop using equal signs!!! ue := for var assignments
	for i := 0; i < 77; i++ {
		fmt.Print(i + 1)
		fmt.Print(": ")

		if i%2 == 0 {
			lit()
		} else {
			tight()
		}

	}

	// made some changes by adding this comment, lets push this up to git too!

	// exit
	tight()
}

func lit() {
	fmt.Println("This is lit!")
}

func tight() {
	fmt.Println("This is tight!")
}

func baii() {
	fmt.Println("Exiting")
}
