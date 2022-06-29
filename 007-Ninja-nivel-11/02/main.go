package main

import (
	"encoding/json"
	"errors"
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

	bs, err := aJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

func aJSON(a person) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("Hubo un erro en aJSON; %v", err))
	}

	return bs, nil
}
