package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user1 struct {
	First string
	Age   int
}

type homie []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user2

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user2

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	fmt.Println("***************************** Exercise 1 *****************************")

	u1 := user1{
		First: "James",
		Age:   32,
	}

	u2 := user1{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user1{
		First: "M",
		Age:   54,
	}

	users1 := []user1{u1, u2, u3}

	fmt.Println(users1)

	// your code goes here
	bs, err := json.Marshal(users1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	fmt.Println("***************************** Exercise 2 *****************************")

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	// declare new homie
	var dudes homie

	// convert "s" to bytestring and unmarshal into dudes
	err = json.Unmarshal([]byte(s), &dudes)
	if err != nil {
		fmt.Println(err)
	}

	// loop and print them all out
	for i, dude := range dudes {
		fmt.Printf("Person: %v \nName: %v %v \nAge: %v\n", i, dude.First, dude.Last, dude.Age)
		fmt.Print("Says: ")
		for _, saying := range dude.Sayings {
			fmt.Printf("\t%v\n", saying)
		}
		fmt.Println()
	}

	fmt.Println("***************************** Exercise 3 *****************************")

	u11 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u22 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u33 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users2 := []user2{u11, u22, u33}

	fmt.Println(users2)

	// your code goes here
	// os.Stdout = the terminal output, can replaced with a file if wanted to (any writer)
	err = json.NewEncoder(os.Stdout).Encode(users2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("***************************** Exercise 4 *****************************")

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("Chaos: ", xi)
	// sort xi (int)
	sort.Ints(xi)
	fmt.Println("Sorted:", xi)

	fmt.Println("Chaos: ", xs)
	// sort xs (string)
	sort.Strings(xs)
	fmt.Println("Sorted:", xs)

	fmt.Println("***************************** Exercise 5 *****************************")

	u111 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u222 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u333 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users3 := []user2{u111, u222, u333}

	fmt.Println("Chaos: ", users3)

	// your code goes here

	// sort by age and print nicely
	sort.Sort(ByAge(users3))
	fmt.Println(">>>> Sort by Age, Sayings Unsorted <<<<")
	for i, user := range users3 {
		fmt.Printf("Person: %v \nName: %v %v \nAge: %v\n", i, user.First, user.Last, user.Age)
		fmt.Print("Says: ")
		for _, saying := range user.Sayings {
			fmt.Printf("\t%v\n", saying)
		}
		fmt.Println()
	}

	// sort by last and print nicely
	sort.Sort(ByLast(users3))
	fmt.Println(">>>> Sort by Last, Sort Sayings Too! <<<<")
	for i, user := range users3 {
		fmt.Printf("Person: %v \nName: %v %v \nAge: %v\n", i, user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		fmt.Print("Says: ")
		for _, saying := range user.Sayings {
			fmt.Printf("\t%v\n", saying)
		}
		fmt.Println()
	}
}
