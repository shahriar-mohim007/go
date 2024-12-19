package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.RWMutex
	counter int
)

func deadlockExample() {
	mu.RLock()         // Acquiring a read lock
	defer mu.RUnlock() // Will release the read lock

	fmt.Println("Read lock acquired")

	mu.Lock()         // Deadlock: Can't acquire write lock while holding a read lock
	defer mu.Unlock() // This line is never reached
	counter++
}
func main() {
	deadlockExample()
}
