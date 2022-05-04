package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas: %v\n", runtime.NumGoroutine())

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		fmt.Println("Entro a la funcion 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Entro a la funcion 2")
		wg.Done()
	}()

	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas: %v\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Funcion Principal")

	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas: %v\n", runtime.NumGoroutine())
}
