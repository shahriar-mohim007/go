package main

import (
	"fmt"
	"sync"
)

var (
	mu1 sync.Mutex
	mu2 sync.Mutex
	wg  sync.WaitGroup
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
	defer wg.Done()
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
	defer wg.Done()

}

func main() {

	wg.Add(4)
	go goroutine1()
	go goroutine2()

	go func() {
		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock() // Goroutine 1 waits for mu2
		defer mu2.Unlock()
		defer wg.Done()
		fmt.Println("first goroutine")
	}()

	go func() {
		mu2.Lock()
		defer mu2.Unlock()

		mu1.Lock() // Goroutine 2 waits for mu1
		defer mu1.Unlock()
		defer wg.Done()
		fmt.Println("second goroutine")
	}()

	// Wait for goroutines to finish
	// (This is not ideal, proper synchronization should be used, but here it's to show the deadlock)
	//var input string
	//fmt.Scanln(&input)
	//fmt.Println(input)
	wg.Wait()
	fmt.Println("all goroutine finished")
}

//A deadlock occurs when two or more goroutines are waiting for each other to release
//resources (in this case, locks), and as a result, none of them can proceed.
//	This can happen when multiple mutexes are acquired in a way that creates a circular dependency between goroutines.
//Example of Deadlock with Mutexes
//
//In the following example, we have two goroutines, each trying to acquire two locks (mutexes) in a different order.
//	This leads to a deadlock because each goroutine is holding one lock and waiting for the other lock, and neither can proceed.
