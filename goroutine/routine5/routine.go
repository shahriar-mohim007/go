package main

import "fmt"

// countTo sends numbers from 0 up to max-1 on a channel
func countTo(max int) <-chan int {
	ch := make(chan int)

	// Start a goroutine to generate numbers
	go func() {
		for i := 0; i < max; i++ {
			ch <- i // Send the current number to the channel
		}
		close(ch) // Close the channel after sending all numbers
	}()

	return ch // Return the channel for the caller to receive from
}

func main() {
	// Range over the values received from the countTo function
	for i := range countTo(10) {
		if i > 5 {
			break
		}
		fmt.Println(i) // Print each number as it is received
	}
}

//The error "fatal error: all goroutines are asleep - deadlock!" occurs because the channel ch is never closed.
//When main tries to range over countTo(10),
//it waits indefinitely for more values, even after all values up to max-1 have been sent.
//
//In Go, when you range over a channel, it must eventually be closed by the sender; otherwise,
//the receiver keeps waiting for more values. To fix this, you need to close the channel ch inside the goroutine
//after sending all the values.
