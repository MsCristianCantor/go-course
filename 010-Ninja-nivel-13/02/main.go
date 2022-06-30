package main

import (
	"fmt"

	"github.com/MsCristianCantor/go-course/010-Ninja-nivel-13/02/cita"
	"github.com/MsCristianCantor/go-course/010-Ninja-nivel-13/02/palabra"
)

func main() {
	fmt.Println(palabra.Conteo(cita.Cuando))

	for k, v := range palabra.ConteoUso(cita.Cuando) {
		fmt.Println(v, k)
	}
}
