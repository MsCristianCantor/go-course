package main

import "github.com/MsCristianCantor/go-course/003-concurrency/uc"

func main() {
	// Basic example of councurrency
	// uc.BasicCouncurrency()

	// Race condition with mutex
	// uc.RaceConditionMutex()

	// Race condition with atomic
	uc.RaceConditionAtomic()
}
