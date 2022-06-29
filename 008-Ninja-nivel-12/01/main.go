package main

import (
	"fmt"

	"github.com/MsCristianCantor/go-course/008-Ninja-nivel-12/01/perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "Chester",
		edad:   perro.Edad(1),
	}
	fmt.Println(c1)
}
