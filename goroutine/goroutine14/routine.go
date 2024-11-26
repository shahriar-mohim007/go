package main

import (
	"fmt"
	"sync"
)

func doThing1() {
	// Simulate some work
	fmt.Println("Doing thing 1")
}

func doThing2() {
	// Simulate some work
	fmt.Println("Doing thing 2")
}

func doThing3() {
	// Simulate some work
	fmt.Println("Doing thing 3")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// Launching goroutines
	go func() {
		defer wg.Done() // Mark this goroutine as done
		doThing1()      // Perform task 1
	}()

	go func() {
		defer wg.Done() // Mark this goroutine as done
		doThing2()      // Perform task 2
	}()

	go func() {
		defer wg.Done() // Mark this goroutine as done
		doThing3()      // Perform task 3
	}()

	// Wait for all goroutines to complete
	wg.Wait()

	// After all goroutines are done
	fmt.Println("All tasks completed.")
}
