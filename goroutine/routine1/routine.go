package main

import (
	"fmt"
)

func process(val int) int {
	// Example processing, e.g., square the value
	return val * val
}

func runThingConcurrently(in <-chan int, out chan<- int) {
	// Start a goroutine to process incoming values concurrently
	go func() {
		for val := range in {
			result := process(val)
			out <- result
		}
		close(out) // Close output channel after processing all input
	}()
}

func main() {
	in := make(chan int)
	out := make(chan int)

	// Run the concurrent processor
	runThingConcurrently(in, out)

	// Send values to the input channel
	go func() {
		for i := 1; i <= 5; i++ {
			in <- i
		}
		close(in) // Close input channel after sending all values
	}()

	// Receive and print results from the output channel
	for result := range out {
		fmt.Println("Processed:", result)
	}
}

//+---------------------+
//|    Main Goroutine   |
//+---------------------+
//| Creates `in` & `out`|
//| channels            |
//|                     |
//| Calls `runThingCon- |
//| currently(in, out)` |
//+----------+----------+
//|
//|
//v
//+---------------------+
//| Processing Goroutine|
//+---------------------+
//| For each val in `in`|
//|   - Process val     |
//|   - Send to `out`   |
//|                     |
//| Closes `out` when   |
//| `in` is closed      |
//+---------------------+
//
//-------------------
//|
//|
//+---------------------+
//|  Sending Goroutine  |
//+---------------------+
//| Sends values 1 - 5  |
//| to `in`             |
//| Closes `in` when all|
//| values are sent     |
//+---------------------+
//
//|
//|
//v
//+---------------------+
//|    Main Goroutine   |
//+---------------------+
//| Reads from `out`    |
//| until `out` is      |
//| closed, printing    |
//| each result         |
//+---------------------+
