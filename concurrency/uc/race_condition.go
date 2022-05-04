package uc

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func RaceConditionMutex() {
	fmt.Println("Número de CPUs", runtime.NumCPU())
	fmt.Println("Número de Gorutinas", runtime.NumGoroutine())
	contador := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			v++
			runtime.Gosched()
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Número de Gorutinas", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta:", contador)
}

func RaceConditionAtomic() {
	fmt.Println("Número de CPUs", runtime.NumCPU())
	fmt.Println("Número de Gorutinas", runtime.NumGoroutine())
	var contador int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Número de Gorutinas", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
