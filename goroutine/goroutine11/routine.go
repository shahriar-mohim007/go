package main

import (
	"fmt"
)

func process(v int) int {
	// Example processing function: just return the value multiplied by 2.
	return v * 2
}

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	var out []int

	// Launch goroutines to process the values from the channel
	for i := 0; i < conc; i++ {
		go func() {
			// Receive value from the input channel and process it
			v := <-ch
			results <- process(v)
		}()
	}

	// Collect results from the results channel
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}

	return out
}

func main() {
	// Create a channel and fill it with some data
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	// Call processChannel to process the values from the channel
	out := processChannel(ch)
	fmt.Println(out)
}
