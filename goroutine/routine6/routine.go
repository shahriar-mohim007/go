package main

import (
	"fmt"
	"sync"
)

// countTo sends numbers from 0 up to max-1 on a channel
func countTo(max int, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)
	wg.Add(1) // Increment WaitGroup counter for this goroutine

	// Start a goroutine to generate numbers
	go func() {
		// Decrement WaitGroup counter when done
		for i := 0; i < max; i++ {
			ch <- i // Send the current number to the channel
		}
		close(ch) // Close the channel after sending all numbers
	}()

	return ch // Return the channel for the caller to receive from
}

func main() {
	var wg sync.WaitGroup
	ch := countTo(10, &wg)

	// Range over the values received from the countTo function
	for i := range ch {
		if i > 5 {
			wg.Done()
			break // Break early if i > 5
		}
		fmt.Println(i) // Print each number as it is received
	}
	// Wait for all goroutines to complete
	wg.Wait() // This will cause a deadlock
}

//Unconsumed Values: The countTo goroutine is still running and attempts to send the remaining
//values (6 to 9) to the channel ch. However,
//since main has stopped receiving,
//these sends block, causing the countTo goroutine to wait indefinitely.
//
//WaitGroup Deadlock: main then calls wg.Wait(), which waits for the countTo goroutine to finish. But since the
//goroutine is blocked trying to send to ch, it never completes and doesn’t call wg.Done(). This causes a deadlock,
//as main waits indefinitely for the countTo goroutine to finish.
