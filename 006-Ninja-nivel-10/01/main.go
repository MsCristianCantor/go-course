package main

import "fmt"

func main() {
	// Solucion  1
	// c := make(chan int)

	// go func() {
	// 	c <- 42
	// }()

	// Solucion  2
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
