package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int)
	in2 := make(chan int)
	done := make(chan struct{})

	// Simulate data sending and closing channels
	go func() {
		for i := 0; i < 5; i++ {
			in <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(in) // Close the first channel
	}()

	go func() {
		for i := 5; i < 10; i++ {
			in2 <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(in2) // Close the second channel
	}()

	// Process incoming data
	for {
		select {
		case v, ok := <-in:
			if !ok {
				in = nil
				if in2 == nil {
					close(done)
				} // Disable this case
				continue
			}
			fmt.Println("Received from in:", v)

		case v, ok := <-in2:
			if !ok {
				in2 = nil // Disable this case
				if in == nil {
					close(done)
				}
				continue
			}
			fmt.Println("Received from in2:", v)

		case <-done:
			fmt.Println("All channels are closed. Exiting.")
			return // Stop when done is closed
		}
	}
}
