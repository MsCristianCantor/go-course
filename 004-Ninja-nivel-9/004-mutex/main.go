package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	con := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := con
			v++
			con = v
			fmt.Println(con)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("valor final: ", con)
}
