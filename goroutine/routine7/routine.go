package main

import (
	"fmt"
	"sync"
)

// countTo sends numbers from 0 up to max-1 on a channel, but stops if it receives a signal from done
func countTo(max int, wg *sync.WaitGroup, done <-chan struct{}) <-chan int {
	ch := make(chan int)
	wg.Add(1) // Increment WaitGroup counter for this goroutine

	// Start a goroutine to generate numbers
	go func() {
		defer wg.Done() // Decrement WaitGroup counter when done
		for i := 0; i < max; i++ {
			select {
			case <-done:
				// Stop if done signal is received
				close(ch)
				return
			case ch <- i:
				// Send the current number to the channel
			}
		}
		close(ch) // Close the channel after sending all numbers
	}()

	return ch // Return the channel for the caller to receive from
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{}) // Create a done channel to signal stopping
	ch := countTo(10, &wg, done)

	// Range over the values received from the countTo function
	for i := range ch {
		if i > 5 {
			close(done) // Signal the goroutine to stop
			break       // Break early if i > 5
		}
		fmt.Println(i) // Print each number as it is received
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

//The done channel pattern provides a way to signal a goroutine that it’s time
//to stop processing. It uses a channel to signal that it’s time to exit.
