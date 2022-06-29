package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name     string
	LastName string
	Phrases  []string
}

func main() {
	p1 := person{
		Name:     "James",
		LastName: "Bond",
		Phrases: []string{
			"Shaken, not stirred",
			"Algun ultimo deseo?",
			"Nunca digas nunca",
		},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
