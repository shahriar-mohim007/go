package main

import (
	"fmt"
	"time"
)

func fanOut(in <-chan int, workerCount int) {
	// Create 'workerCount' number of workers
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			for num := range in {
				// Simulate some processing
				fmt.Printf("Worker %d processed %d\n", workerID, num)
				time.Sleep(100 * time.Millisecond) // Simulate work delay
			}
		}(i)
	}
}

func main() {
	// Create a channel to send data to workers
	input := make(chan int)

	// Number of workers
	workerCount := 3

	// Start workers
	fanOut(input, workerCount)

	// Send numbers to the input channel
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input) // Close the input channel when done
	}()

	// Allow time for all workers to finish processing
	time.Sleep(2 * time.Second)
	fmt.Println("All work done.")
}
