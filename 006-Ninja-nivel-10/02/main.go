package main

import "fmt"

func main() {
	// EJERCICIO 1
	// cs := make(chan<- int)
	// Solucion hacemos el canal bidireccional
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	// EJERCICIO 2
	// cr := make(<-chan int)
	// Solucion hacemos el canal bidireccional
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cr)
}
