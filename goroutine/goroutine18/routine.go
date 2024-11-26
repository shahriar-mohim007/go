package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.RWMutex

func read() int {
	mu.RLock()         // Acquire read lock
	defer mu.RUnlock() // Release read lock
	return counter
}

func write(value int) {
	mu.Lock()         // Acquire write lock
	defer mu.Unlock() // Release write lock
	counter = value
}

func main() {
	var wg sync.WaitGroup
	// Simulate concurrent reads and writes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			write(i)
			fmt.Println("Written:", i)
		}(i)
	}

	// Concurrently read the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Read:", read())
		}()
	}

	wg.Wait()
}

//sync.RWMutex (Read-Write Mutex)
//
//This type of mutex allows multiple goroutines to read a shared resource concurrently while ensuring that only one goroutine can write to it at a time.
//It is useful when you have many goroutines reading shared data, but only a few need to modify it, as it allows for better concurrency than a regular Mutex.
//
//Methods:
//
//Lock(): Acquires a write lock. Only one goroutine can hold the write lock at a time, and no other goroutine can acquire a read or write lock while it's held.
//Unlock(): Releases the write lock.
//RLock(): Acquires a read lock. Multiple goroutines can acquire the read lock concurrently as long as no goroutine has the write lock.
//RUnlock(): Releases the read lock.
