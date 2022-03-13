package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// struct for loading JSON
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   uint   `json:"Age"`
}

func main() {
	// ******** Marshalling ********
	// create new person
	p1 := person{
		First: "Jett",
		Last:  "Krassner",
		Age:   25,
	}

	p2 := person{
		First: "Jelly",
		Last:  "Krassner",
		Age:   22,
	}

	people := []person{p1, p2}

	// marshal into json string (byte-slice)
	bs, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	// print json string
	fmt.Println(bs)

	// ******** Unmarshalling ********
	jsonString := `[{"First":"Jett","Last":"Krassner","Age":25},{"First":"Jelly","Last":"Krassner","Age":22}]`

	// convert to bs
	bs = []byte(jsonString)

	// create new var type struct for loading
	var jsonPeople []person

	// unmarshal (the pointer to the person type)
	err = json.Unmarshal(bs, &jsonPeople)
	if err != nil {
		log.Fatal(err)
	}

	// print the loaded struct
	fmt.Println(jsonPeople)

	// ******** Sorting ********
	
}
