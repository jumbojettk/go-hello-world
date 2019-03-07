package main

// import pkgs
import "fmt"

func main() {
	fmt.Println("Uh Suhhh!")
	fmt.Println("I'm Jett, and I'm", 25, "years old!")
	lit()

	fmt.Println("Bouta print \"lit\" during \"even\" and running 77 times!")

	// notice the parenthesis not used, no need
	// stop using equal signs!!! ue := for var assignments
	for i := 0; i < 7; i++ {
		fmt.Print(i+1, ": ") // use comma , for printing different types
		if i%2 == 0 {
			lit()
		} else {
			tight()
		}

	}

	// made some changes by adding this comment, lets push this up to git too!
	// _ ignores, so this ignores the error
	n, _ := fmt.Print("Let's see how many bytes this sentence is: ")
	fmt.Println(n, "!")

	// declare a new var with :=
	cool := 10
	fmt.Println("Cool is:", cool)
	cool = 20 // re assignment with =
	fmt.Println("Cool with new assignment is:", cool)

	// print exit
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
