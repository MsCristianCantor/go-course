package main

import (
	"fmt"

	"github.com/MsCristianCantor/go-course/010-Ninja-nivel-13/01/codigo-inicial/perrito"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	fido := canino{
		nombre: "Fido",
		edad:   perrito.Edad(10),
	}
	fmt.Println(fido)
	fmt.Println(perrito.EdadDos(20))
}
