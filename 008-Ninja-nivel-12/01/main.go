package main

import (
	"fmt"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "Chester",
		// edad:   perro.Edad(2),
	}
	fmt.Println(c1)
}
