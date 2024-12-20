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
	counter += value
}

func main() {
	var wg sync.WaitGroup
	// Simulate concurrent reads and writes
	wg.Add(6)
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		write(i)
	//		fmt.Println("Written:", i)
	//	}(i)
	//}
	//write(10)
	go func(i int) {
		defer wg.Done()
		write(i)
		fmt.Println("Written:", counter)
	}(1)

	// Concurrently read the counter
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		fmt.Println("Read:", read())
	//	}()
	//}
	go func() {
		defer wg.Done()
		fmt.Println("Read:", read())
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Read:", read())
	}()
	go func(i int) {
		defer wg.Done()
		write(i)
		fmt.Println("Written:", counter)
	}(1)
	//fmt.Println("Main goroutine", read())
	go func(i int) {
		defer wg.Done()
		write(i)
		fmt.Println("Written:", counter)
	}(1)

	go func() {
		defer wg.Done()
		fmt.Println("Read:", read())
	}()

	wg.Wait()
	fmt.Println("All goroutines finished", read())
}
