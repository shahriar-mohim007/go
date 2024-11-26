package main

import (
	"fmt"
	"sync"
)

var (
	mu1 sync.Mutex
	mu2 sync.Mutex
)

func goroutine1() {
	mu1.Lock()
	fmt.Println("Goroutine 1 acquired mu1")

	// Simulating some work
	// Now goroutine1 tries to acquire mu2
	mu2.Lock()
	fmt.Println("Goroutine 1 acquired mu2")

	// Do some work with both locks
	mu2.Unlock()
	mu1.Unlock()
}

func goroutine2() {
	mu2.Lock()
	fmt.Println("Goroutine 2 acquired mu2")

	// Simulating some work
	// Now goroutine2 tries to acquire mu1
	mu1.Lock()
	fmt.Println("Goroutine 2 acquired mu1")

	// Do some work with both locks
	mu1.Unlock()
	mu2.Unlock()
}

func main() {
	go goroutine1()
	go goroutine2()

	// Wait for goroutines to finish
	// (This is not ideal, proper synchronization should be used, but here it's to show the deadlock)
	var input string
	fmt.Scanln(&input)
}

//A deadlock occurs when two or more goroutines are waiting for each other to release resources (in this case, locks), and as a result, none of them can proceed.
//	This can happen when multiple mutexes are acquired in a way that creates a circular dependency between goroutines.
//Example of Deadlock with Mutexes
//
//In the following example, we have two goroutines, each trying to acquire two locks (mutexes) in a different order.
//	This leads to a deadlock because each goroutine is holding one lock and waiting for the other lock, and neither can proceed.
