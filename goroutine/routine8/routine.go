package main

import (
	"fmt"
	"time"
)

// searchData performs searches concurrently using multiple searchers.
// It returns the first result it receives and stops the other searchers.
func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})   // Channel to signal other goroutines to stop
	result := make(chan []string) // Channel to receive results

	// Launch each searcher in a separate goroutine
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s): // Send search result if successful
			case <-done: // Exit if another result has been received
			}
		}(searcher)
	}

	// Wait for the first result from any searcher
	r := <-result
	close(done) // Close done channel to signal other goroutines to exit

	return r
}

// Example search functions
func search1(s string) []string {
	time.Sleep(1 * time.Millisecond) // Simulate delay
	return []string{"Result from search1 with " + s}
}

func search2(s string) []string {
	time.Sleep(500 * time.Millisecond) // Simulate faster delay
	return []string{"Result from search2 with " + s}
}

func main() {
	// Define search functions
	searchers := []func(string) []string{search1, search2}

	// Run search and print the first result
	result := searchData("example search term", searchers)
	fmt.Println("First result:", result)
}

//+-----------------+      +----------------+      +----------------+
//|  Main Function  |      |  Searcher 1    |      |  Searcher 2    |
//|-----------------|      |----------------|      |----------------|
//|                 |      |                |      |                |
//|  done := make(chan struct{})                            |
//|  result := make(chan []string)                          |
//|                                                         |
//|  for _, searcher := range searchers {                   |
//|      go func(searcher) {                                |
//|          select {                                       |
//|            case result <- searcher(s):   <-----------+  |
//|            case <-done:  (exit)                       |  |
//|          }                                            |  |
//|      }(searcher)                                      |  |
//|  }                                                    |  |
//|                   |                                    |  |
//|                   +------------------------------------+  |
//|                   |  result <- (result from any searcher) |
//|                   +------------------------------------>  |
//|                   |                                    |  |
//|                   |<-done  (notify remaining searchers)   |
//|                   |                                    |  |
//+-------------------+------------------------------------+--+
//
//
//Goroutines Start: Each searcher function starts as a goroutine and runs the select statement.
//Waiting on Result: Each searcher either sends a result to result or listens for a signal on done to exit.
//Result Found: The first searcher to finish sends its result to result, which is received by the main function.
//Signal Others to Exit: After receiving the first result, the main function closes done, signaling all other goroutines to exit without further processing.
