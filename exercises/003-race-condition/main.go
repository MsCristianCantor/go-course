package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	con := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := con
			runtime.Gosched()
			v++
			con = v
			fmt.Println(con)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("valor final: ", con)
}
