package main

import "concurrency/uc"

func main() {
	// Basic example of councurrency
	// uc.BasicCouncurrency()

	// Race condition with mutex
	// uc.RaceConditionMutex()

	// Race condition with atomic
	uc.RaceConditionAtomic()
}
