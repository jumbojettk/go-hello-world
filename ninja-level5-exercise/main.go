package main

import (
	"fmt"
)

// EX 1: Create your own type “person”
// which will have an underlying type of “struct”
// so that it can store the following data: first name, last name, favorite ice cream flavors
type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

// EX 2: Create a new type: vehicle.
// The underlying type is a struct.
// The fields: doors, color
type vehicle struct {
	doors int
	color string
}

// Create two new types: truck & sedan.
// The underlying type of each of these new types is a struct.
// Embed the “vehicle” type in both truck & sedan.
// Give truck the field “fourWheel” which will be set to bool.
// Give sedan the field “luxury” which will be set to bool.
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println("******************** EXERCISE 1 ********************")
	// Create two VALUES of TYPE person.
	p1 := person{
		firstName:   "Jett",
		lastName:    "Krassnerski",
		favIceCream: []string{"Butter pecan", "Chocolate chip cookie dough", "Vanilla"},
	}

	p2 := person{
		firstName:   "Jelly",
		lastName:    "Krassner",
		favIceCream: []string{"Chocolate chip cookie dough", "Chocolate chucks", "Strawberry"},
	}

	// Print out the values, ranging over the elements in the slice which stores the favorite flavors.
	fmt.Println("Name:", p1.firstName, p1.lastName)
	fmt.Print("Favorite Ice Cream Flavs: ")
	for i, v := range p1.favIceCream {
		fmt.Print(i, v)
	}
	fmt.Println()

	fmt.Println("Name:", p2.firstName, p2.lastName)
	fmt.Print("Favorite Ice Cream Flavs: ")
	for i, v := range p2.favIceCream {
		fmt.Print(i, v)
	}
	fmt.Println()

	fmt.Println("******************** EXERCISE 2 ********************")
	// Take the code from the previous exercise,
	// then store the values of type person in a map with the key of last name.
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	// Access each value in the map.
	// Print out the values, ranging over the slice.
	for _, p := range m {
		fmt.Println("Name:", p.firstName, p.lastName)
		fmt.Print("Favorite Ice Cream Flavs: ")
		for _, favIce := range p.favIceCream {
			fmt.Print(favIce)
			fmt.Print(", ")
		}
		fmt.Println()
	}

	fmt.Println("******************** EXERCISE 3 ********************")
	// Using the vehicle, truck, and sedan structs:
	// using a composite literal, create a value of type truck and assign values to the fields
	monster := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	// using a composite literal, create a value of type sedan and assign values to the fields.
	m5 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	// Print out each of these values.
	fmt.Printf("Monster Truck: %v\nM5 Sedan: %v\n", monster, m5)

	// Print out a single field from each of these values.
	fmt.Println("Monster Truck")
	fmt.Printf("Type: %T \t Doors: %v \t Color: %v \t Four wheel: %v\n", monster, monster.doors, monster.color, monster.fourWheel)

	fmt.Println("BMW M5 Sedan")
	fmt.Printf("Type: %T \t Doors: %v \t Color: %v \t Four wheel: %v\n", m5, m5.doors, m5.color, m5.luxury)

	fmt.Println("******************** EXERCISE 4 ********************")
	// Create and use an anonymous struct.
	fam := struct {
		lastName  string
		members   []string
		memberAge map[string]int
		tight     bool
	}{
		lastName: "Krassner",
		members:  []string{"Jett", "Jelly", "Frank", "Piyawadee"},
		memberAge: map[string]int{
			"Jett":      25,
			"Jelly":     22,
			"Frank":     70,
			"Piyawadee": 50,
		},
		tight: true,
	}

	// print them out
	fmt.Println("Family:", fam.lastName)

	// range all members
	fmt.Print("Members: ")
	for _, mem := range fam.members {
		fmt.Print(mem, ", ")
	}
	fmt.Println()

	// range all members and their age
	fmt.Print("Members' age: ")
	for key, age := range fam.memberAge {
		fmt.Print(key, " is ", age, ", ")
	}
	fmt.Println()

}
