package main

// import pkgs
import "fmt"

// declaring with var and default values
var globalDec = 7
var globalDecDefaultInt int
var globalDecDefaultStr string

// declaring new type
type dope int

var nice dope

func main() {
	fmt.Println("Uh Suhhh!")
	fmt.Println("I'm Jett, and I'm", 25, "years old!")
	lit()

	// exploring declaration with var
	fmt.Println("The gobalDec var is:", globalDec)
	fmt.Println("The gobalDecDefaultInt without assignment is:", globalDecDefaultInt)
	fmt.Println("The gobalDecDefaultStr without assignment is:", globalDecDefaultStr, "\t(should be an empty string)")

	// exploring Printf, template string
	// %T is type of the variable and \n for newline
	globalDecDefaultInt = 77
	globalDecDefaultStr = "Everything is awesome! Every situation is a lituation!"
	fmt.Printf("globalDecInt, value is %v, type is: %T\n", globalDecDefaultInt, globalDecDefaultInt)
	fmt.Printf("globalDecStr, value is %v, type is: %T\n", globalDecDefaultStr, globalDecDefaultStr)

	// exploring type
	nice = 10
	fmt.Printf("The variable nice has value of %v, and type of %T\n", nice, nice)

	// since globalDec = nice will not work due to diff types
	// converting the type "dope" to "int" (not called casting)
	globalDec = int(nice)
	fmt.Printf("The variable gobalDec should now have value of nice, has value of %v, and type of %T\n", globalDec, globalDec)

	// using string literals
	fmt.Println(`This is a literal string with "another string inside with quotes" it even prints \t and \n`)

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

	// declare a new var and assign with :=
	cool := 10
	fmt.Println("Cool is:", cool)
	cool = 20 // re assignment with =
	fmt.Println("Cool with new assignment is:", cool)

	// print exit
	baii()
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
