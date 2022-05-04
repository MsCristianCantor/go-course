package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Printf("Hola soy %s \n", p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		nombre: "pepito",
	}

	// diAlgo(p1)
	// diAlgo(&p1)
	p1.hablar()
}
