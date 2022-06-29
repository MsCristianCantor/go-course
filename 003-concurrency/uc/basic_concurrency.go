package uc

import (
	"fmt"
	"runtime"
	"sync"
)

var wgb sync.WaitGroup

// BasicCouncurrency to test concurrency
func BasicCouncurrency() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	wgb.Add(1)

	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	wgb.Wait()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	wgb.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
