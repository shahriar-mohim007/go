package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

// Increment function uses mutex to safely modify the shared 'counter'
func increment() {
	mutex.Lock()         // Lock the mutex to ensure only one goroutine can access the critical section
	defer mutex.Unlock() // Unlock the mutex when done

	counter++ // Critical section
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple goroutines to increment the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait() // Wait for all goroutines to complete

	fmt.Println("Final counter value:", counter) // Prints 5
}

//In Go, a mutex (short for mutual exclusion) is used to synchronize access to shared resources among multiple goroutines.
//	When multiple goroutines are trying to read or modify the same data concurrently, it can lead to race conditions, which can cause unpredictable behavior.
//	A mutex provides a way to ensure that only one goroutine can access the shared resource at a time, preventing concurrent access issues.
//
//Go provides the sync.Mutex type, which is part of the sync package, to achieve mutual exclusion.
