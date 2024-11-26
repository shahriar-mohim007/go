package main

import (
	"fmt"
	"time"
)

type Searcher struct {
	name string
}

func (s *Searcher) Search(result chan<- string, done <-chan struct{}) {
	// Simulate searching
	time.Sleep(time.Second)
	// Send the result to the result channel
	select {
	case result <- s.name:
	case <-done:
		fmt.Printf("%s stopped\n", s.name)

	}
}

func main() {
	done := make(chan struct{}) // Channel to signal goroutines to stop
	result := make(chan string) // Channel to get search results

	searchers := []Searcher{
		{name: "Searcher 1"},
		{name: "Searcher 2"},
		{name: "Searcher 3"},
	}

	// Launch a goroutine for each searcher
	for _, s := range searchers {
		go s.Search(result, done)
	}

	// Wait for the first result and then stop all goroutines
	select {
	case res := <-result:
		fmt.Println("First result:", res)
		// Close the done channel to signal all goroutines to stop
		close(done)
	}
	time.Sleep(time.Second)
}

//The Done Channel Pattern
