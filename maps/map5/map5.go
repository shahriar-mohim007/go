package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := make(map[string]int)
	var mu sync.RWMutex

	// Writer goroutine 1
	go func() {
		mu.Lock()
		data["key"] = 100
		mu.Unlock()
	}()

	// Reader goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			mu.RLock()
			fmt.Println("Read:", data["key"])
			mu.RUnlock()
		}
	}()

	// Writer goroutine 2
	go func() {
		mu.Lock()
		data["key"] = 10
		mu.Unlock()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Final value:", data["key"])
}
